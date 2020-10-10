package redis

import (
	"github.com/firmeve/firmeve/kernel"
)

type Provider struct {
	kernel.BaseProvider
}

func (Provider) Name() string {
	return `redis`
}

func (p *Provider) Register() {
	redisConfig := new(Configuration)
	p.Config.Bind(`redis`, redisConfig)
	redis := New(redisConfig)
	p.Bind(`redis`, redis)
}

func (Provider) Boot() {

}
