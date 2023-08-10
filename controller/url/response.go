package url_controller

import "url-shortener/service/url"

type CreateBulkResponse struct {
	ShortUrl string
	LongUrl  string
}

func ToCreateBulkResponse(items []url_service.CreateUrlItem) []CreateBulkResponse {
	results := make([]CreateBulkResponse, len(items))
	for i, item := range items {
		results[i] = CreateBulkResponse{
			ShortUrl: item.ShortUrl,
			LongUrl:  item.LongUrl,
		}
	}
	return results
}
