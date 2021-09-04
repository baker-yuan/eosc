/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package traffic

var (
	controller = NewController()
)

//Listener 设置默认监听器
func Listener(network string, addr string) error {
	return controller.Listener(network, addr)
}

func All() Traffics {
	return controller.All()
}

func Close() {
	controller.Close()
}