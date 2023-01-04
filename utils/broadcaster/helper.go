package broadcast

var broadcasterSingleton Broadcaster

func GetBroadcaster() Broadcaster {
	if broadcasterSingleton == nil {
		broadcasterSingleton = NewBroadcaster(1)
	}

	return broadcasterSingleton
}
