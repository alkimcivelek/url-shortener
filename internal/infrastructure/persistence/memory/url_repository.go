package memory

import (
	"sync"
	"url-shortener/internal/domain/entity"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/valueobject"
)

type URLRepositoryMemory struct {
	urls           map[string]*entity.URL
	urlsByOriginal map[string]*entity.URL
	mutex          sync.RWMutex
}

func NewURLRepositoryMemory() repository.URLRepository {
	return &URLRepositoryMemory{
		urls:           make(map[string]*entity.URL),
		urlsByOriginal: make(map[string]*entity.URL),
	}
}

func (r *URLRepositoryMemory) Save(url *entity.URL) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	originalURLStr := url.OriginalURL().Value()
	shortCodeStr := url.ShortCode().Value()

	// Check duplicates
	if _, exists := r.urlsByOriginal[originalURLStr]; exists {
		return repository.ErrDuplicateURL
	}

	if _, exists := r.urls[shortCodeStr]; exists {
		return repository.ErrDuplicateURL
	}

	r.urls[shortCodeStr] = url
	r.urlsByOriginal[originalURLStr] = url
	return nil
}

func (r *URLRepositoryMemory) FindByShortCode(shortCode valueobject.ShortCode) (*entity.URL, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	url, exists := r.urls[shortCode.Value()]
	if !exists {
		return nil, repository.ErrURLNotFound
	}

	return url, nil
}

func (r *URLRepositoryMemory) FindByOriginalURL(originalURL valueobject.OriginalURL) (*entity.URL, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	url, exists := r.urlsByOriginal[originalURL.Value()]
	if !exists {
		return nil, repository.ErrURLNotFound
	}

	return url, nil
}

func (r *URLRepositoryMemory) Update(url *entity.URL) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	shortCodeStr := url.ShortCode().Value()

	if _, exists := r.urls[shortCodeStr]; !exists {
		return repository.ErrURLNotFound
	}

	r.urls[shortCodeStr] = url
	r.urlsByOriginal[url.OriginalURL().Value()] = url
	return nil
}

func (r *URLRepositoryMemory) GetAll() ([]*entity.URL, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	urls := make([]*entity.URL, 0, len(r.urls))
	for _, url := range r.urls {
		urls = append(urls, url)
	}

	return urls, nil
}
