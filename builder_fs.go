package avatar

import (
	"strings"

	"github.com/go-universal/cache"
	"github.com/go-universal/utils"
)

func (b *builder) WithStorage(root string) FactoryBuilder {
	b.factory.storage = utils.NormalizePath(root)
	return b
}

func (b *builder) WithPrefix(prefix string) FactoryBuilder {
	b.factory.prefix = strings.TrimSpace(prefix)
	return b
}

func (b *builder) WithQueue(queue cache.Queue) FactoryBuilder {
	b.factory.queue = queue
	return b
}

func (b *builder) WithNumberedFile() FactoryBuilder {
	b.factory.numbered = true
	return b
}

func (b *builder) WithTimestampedFile() FactoryBuilder {
	b.factory.numbered = false
	return b
}
