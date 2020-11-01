package main

import (
	"strconv"
	"strings"
)

// StorageMemory ...
type StorageMemory struct {
	products []*Product
}

// Find ...
func (s *StorageMemory) Find(criteria string) ([]*Product, error) {
	id, err := strconv.Atoi(criteria)
	if err != nil {
		return s.findByText(criteria)
	}

	return s.findByID(id)
}

func (s *StorageMemory) findByID(criteria int) ([]*Product, error) {
	var ps []*Product

	for _, p := range getProducts() {
		if p.ID == criteria {
			ps = append(ps, &p)
			break
		}
	}

	return ps, nil
}

func (s *StorageMemory) findByText(criteria string) ([]*Product, error) {
	var ps []*Product

	for _, p := range getProducts() {
		if strings.Contains(p.Brand, criteria) || strings.Contains(p.Description, criteria) {
			ps = append(ps, &p)
		}
	}

	return ps, nil
}

func getProducts() []Product {
	return []Product{
		{ID: 3, Brand: "weñxoab", Description: "hqhoy qacirk", Image: "www.lider.cl/catalogo/images/homeIcon.svg", Price: 171740},
		{ID: 86, Brand: "tsf pzlflrg", Description: "hqhoy qacirk", Image: "www.lider.cl/catalogo/images/whiteLineIcon.svg", Price: 91993},
		{ID: 281, Brand: "tsf pzlflrg", Description: "vhxpfo oxsnm", Image: "www.lider.cl/catalogo/images/bedRoomIcon.svg", Price: 648166},
		{ID: 12, Brand: "vfbjgpt", Description: "iwpazñ ltxsh", Image: "www.lider.cl/catalogo/images/tvIcon.svg", Price: 647307},
		{ID: 121, Brand: "erehzgj", Description: "gzifl ngfxpr", Image: "www.lider.cl/catalogo/images/electronicsIcon.svg", Price: 426816},
		{ID: 123, Brand: "weñxoab", Description: "ljñkv bmfwuq", Image: "www.lider.cl/catalogo/images/furnitureIcon.svg", Price: 581042},
	}
}
