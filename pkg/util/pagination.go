package util

type pageMeta struct {
	Page   uint32
	Limit  uint32
	Offset uint32
}

func PageMeta(page, limit uint32) pageMeta {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := limit * (page - 1)
	pageMeta := pageMeta{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
	return pageMeta
}
