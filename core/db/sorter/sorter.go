package sorter

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SortItem struct {
	Field    string `json:"field"`
	SortType string `json:"sortType"`
}

func parseSort(sort string) (itemList []SortItem) {
	sort = strings.TrimSpace(sort)
	if sort == "" {
		return
	}
	sortList := strings.Split(sort, ",")
	for _, sStr := range sortList {
		sortStr := strings.TrimSpace(sStr)
		if sStr == "" {
			continue
		}
		sortItem := SortItem{}
		if strings.HasPrefix(sortStr, "-") {
			sortItem.Field = strings.TrimPrefix(sortStr, "-")
			sortItem.SortType = "desc"
		} else {
			sortItem.Field = sortStr
			sortItem.SortType = "asc"
		}
		itemList = append(itemList, sortItem)
	}
	return itemList
}

// Sort sql排序处理
// sort: "id,-name" "id asc, name desc"
func Sort(tx *gorm.DB, sort string) {
	sortList := parseSort(sort)
	for _, item := range sortList {
		tx.Order(fmt.Sprintf("%s %s", item.Field, item.SortType))
	}
}
