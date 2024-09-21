package eventbus

import (
	"reflect"
	"sync"
	"testing"
)

func TestEventBus_HasCallback(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if got := bus.HasCallback(tt.args.topic); got != tt.want {
				t.Errorf("HasCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_Publish(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			bus.Publish(tt.args.topic, tt.args.args...)
		})
	}
}

func TestEventBus_Subscribe(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.Subscribe(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeAsync(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic         string
		fn            interface{}
		transactional bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.SubscribeAsync(tt.args.topic, tt.args.fn, tt.args.transactional); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeAsync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeOnce(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.SubscribeOnce(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeOnce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeOnceAsync(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.SubscribeOnceAsync(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeOnceAsync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_Unsubscribe(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic   string
		handler interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.Unsubscribe(tt.args.topic, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_WaitAsync(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			bus.WaitAsync()
		})
	}
}

func TestEventBus_doPublish(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		handler *eventHandler
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			bus.doPublish(tt.args.handler, tt.args.args...)
		})
	}
}

func TestEventBus_doPublishAsync(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		handler *eventHandler
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			bus.doPublishAsync(tt.args.handler, tt.args.args...)
		})
	}
}

func TestEventBus_doSubscribe(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic   string
		fn      interface{}
		handler *eventHandler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if err := bus.doSubscribe(tt.args.topic, tt.args.fn, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("doSubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_findHandlerIdx(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic    string
		callback reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if got := bus.findHandlerIdx(tt.args.topic, tt.args.callback); got != tt.want {
				t.Errorf("findHandlerIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_removeHandler(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		topic string
		idx   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			bus.removeHandler(tt.args.topic, tt.args.idx)
		})
	}
}

func TestEventBus_setUpPublish(t *testing.T) {
	type fields struct {
		handlers map[string][]*eventHandler
		lock     sync.Mutex
		wg       sync.WaitGroup
	}
	type args struct {
		callback *eventHandler
		args     []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := &EventBus{
				handlers: tt.fields.handlers,
				lock:     tt.fields.lock,
				wg:       tt.fields.wg,
			}
			if got := bus.setUpPublish(tt.args.callback, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setUpPublish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Bus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
