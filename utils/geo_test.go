package utils

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type dis struct {
	distance float64
}

type dis1 []dis

func (a1 dis1) Len() int { // 重写 Len() 方法
	return len(a1)
}
func (a1 dis1) Swap(i, j int) { // 重写 Swap() 方法
	a1[i], a1[j] = a1[j], a1[i]
}
func (a1 dis1) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a1[j].distance > a1[i].distance
}

func TestGeo(t *testing.T) {
	t.Log(" start ", time.Now())

	var a []dis

	lat1 := 23.1378010917
	lng1 := 113.4022203113
	lat2 := 22.1191433172
	lng2 := 113.5826193044
	a = append(a, dis{distance: GeoDistance(lng1, lat1, lng2, lat2)})
	a = append(a, dis{distance: GeoDistance(13.361389, 38.115556, 15.087269, 37.502669)})
	a = append(a, dis{distance: GeoDistance(121.596829565, 28.289153604, 121.59682534207816, 28.289150365195056)})
	a = append(a, dis{distance: GeoDistance(121.4662435, 31.28984375, 121.4665202, 31.28952637)})

	sort.Sort(dis1(a))

	fmt.Println(a)
	t.Log(" end ", time.Now())
}
