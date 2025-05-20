package router

import "reflect"

// Reflection-based approach
func isStructEmpty(v interface{}) bool {
    return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
