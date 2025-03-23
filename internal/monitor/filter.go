package monitor

func FilterRestartedPods(pods []PodRestartInfo, minRestarts int32) []PodRestartInfo {
	var filtered []PodRestartInfo
	for _, pod := range pods {
		if pod.Restarts >= minRestarts {
			filtered = append(filtered, pod)
		}
	}
	return filtered
}
