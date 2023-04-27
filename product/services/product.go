package services

import (
	"context"
	productPb "learning-grpc/product/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(ctx context.Context, pageParam *productPb.Page) (*productPb.Products, error) {
	var products []*productPb.Product
	var pagination productPb.Pagination

	var page, total, limit, offset int64 = 1, 1, 1, 1

	if pageParam.GetPage() != 0 {
		page = pageParam.GetPage()
	}

	sql := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name")

	sql.Count(&total)

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagination.Total = uint64(total)
	pagination.PerPage = uint64(limit)
	pagination.CurrentPage = uint32(page)
	pagination.LastPage = uint32(math.Ceil(float64(total) / float64(limit)))

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(
			codes.Internal,
			err.Error(),
		)
	}
	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Error featching data %v\n", err.Error())
		}
		product.Category = &category
		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return response, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *productPb.Id) (*productPb.Product, error) {
	row := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name").
		Where("p.id = ?", id.GetId()).Row()

	var product productPb.Product
	var category productPb.Category

	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
		log.Printf("Error featching data %v\n", err.Error())
		return nil, err
	}

	product.Category = &category

	return &product, nil
}

// func (p *ProductService) CreateProduct(context.Context, *productPb.Product) (*productPb.Id, error) {

// }

// UpdateProduct(context.Context, *Product) (*Status, error)
// DeleteProduct(context.Context, *Id) (*Status, error)
