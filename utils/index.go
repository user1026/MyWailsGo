package utils

func BToGb(num uint64) float64 {
	var Gb uint64 = 1024 * 1024 * 1024
	var Mb uint64 = 1024 * 1024
	if (num / Gb) > 1 {
		return float64(num) / float64(Gb)
	} else if (num / Mb) > 1 {
		return float64(num) / float64(Mb)
	}
	return 0.00
}
