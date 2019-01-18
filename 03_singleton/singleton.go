package singleton

import "sync"

// 单例类
type Singleton struct{}

var (
	s1, s2, s3, s4 *Singleton
	lock1, lock2   sync.Mutex
	once           sync.Once
)

// 这是我们最容易想到也是最简单的单例模式。在单线程中完全正确，但是在并发时，就可能出现问题。
// 例如：
// 协成A执行到s1 = &Singleton{}之前，协程B也来获取实例，由于协成A还没有构建实例，协成B也会执行s1 = &Singleton{}，故实际不是单例。
func GetInstance1() *Singleton {
	if s1 == nil {
		s1 = &Singleton{}
	}
	return s1
}

// 为了解决并发的问题，我们采用锁机制，这样就能保证在并发情况下，只有一个协成获取实例。
// 每次调用都加锁，但是加锁的成本是很高的，有没有优化的空间呢？
func GetInstance2() *Singleton {
	lock1.Lock()
	defer lock1.Unlock()
	if s2 == nil {
		s2 = &Singleton{}
	}
	return s2
}

// 在实例为空时才加锁，避免每次调用都加锁，提高了代码的执行效率。
func GetInstance3() *Singleton {
	if s3 == nil {
		lock2.Lock()
		defer lock2.Unlock()
		if s3 == nil {
			s3 = &Singleton{}
		}
	}
	return s3
}

// 但是go中还有更优雅的方式。
// 回想单例模式概念，就是只有一个单例。那么我让生成单例的代码只执行一次不就行了吗。
func GetInstance4() *Singleton {
	once.Do(func() {
		s4 = &Singleton{}
	})
	return s4
}
