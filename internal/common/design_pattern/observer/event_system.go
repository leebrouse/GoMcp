package observer

import (
	"fmt"
	"sync"
	"time"
)

// EventType 事件类型
type EventType string

const (
	EventToolExecuted  EventType = "tool_executed"
	EventErrorOccurred EventType = "error_occurred"
	EventSystemStart   EventType = "system_start"
	EventSystemStop    EventType = "system_stop"
)

// Event 事件结构
type Event struct {
	Type      EventType              `json:"type"`
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
	Source    string                 `json:"source"`
}

// Observer 观察者接口
type Observer interface {
	OnEvent(event Event)
	GetID() string
}

// Subject 主题接口
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(event Event)
}

// EventBus 事件总线
type EventBus struct {
	observers map[EventType][]Observer
	mu        sync.RWMutex
}

// NewEventBus 创建事件总线
func NewEventBus() *EventBus {
	return &EventBus{
		observers: make(map[EventType][]Observer),
	}
}

// Attach 添加观察者
func (eb *EventBus) Attach(eventType EventType, observer Observer) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if eb.observers[eventType] == nil {
		eb.observers[eventType] = make([]Observer, 0)
	}
	eb.observers[eventType] = append(eb.observers[eventType], observer)
}

// Detach 移除观察者
func (eb *EventBus) Detach(eventType EventType, observer Observer) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	observers := eb.observers[eventType]
	for i, obs := range observers {
		if obs.GetID() == observer.GetID() {
			eb.observers[eventType] = append(observers[:i], observers[i+1:]...)
			break
		}
	}
}

// Notify 通知观察者
func (eb *EventBus) Notify(event Event) {
	eb.mu.RLock()
	observers := eb.observers[event.Type]
	eb.mu.RUnlock()

	for _, observer := range observers {
		go observer.OnEvent(event)
	}
}

// Publish 发布事件
func (eb *EventBus) Publish(eventType EventType, source string, data map[string]interface{}) {
	event := Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
		Source:    source,
	}
	eb.Notify(event)
}

// LoggerObserver 日志观察者
type LoggerObserver struct {
	id string
}

func NewLoggerObserver(id string) *LoggerObserver {
	return &LoggerObserver{id: id}
}

func (l *LoggerObserver) OnEvent(event Event) {
	fmt.Printf("[%s] %s - %s: %+v\n",
		event.Timestamp.Format("2006-01-02 15:04:05"),
		event.Type,
		event.Source,
		event.Data)
}

func (l *LoggerObserver) GetID() string {
	return l.id
}

// MetricsObserver 指标观察者
type MetricsObserver struct {
	id      string
	metrics map[string]int64
	mu      sync.RWMutex
}

func NewMetricsObserver(id string) *MetricsObserver {
	return &MetricsObserver{
		id:      id,
		metrics: make(map[string]int64),
	}
}

func (m *MetricsObserver) OnEvent(event Event) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := string(event.Type)
	m.metrics[key]++

	// 可以在这里添加指标上报逻辑
	fmt.Printf("Metrics [%s]: %s = %d\n", m.id, key, m.metrics[key])
}

func (m *MetricsObserver) GetID() string {
	return m.id
}

func (m *MetricsObserver) GetMetrics() map[string]int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make(map[string]int64)
	for k, v := range m.metrics {
		result[k] = v
	}
	return result
}
