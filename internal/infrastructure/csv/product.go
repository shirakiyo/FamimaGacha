package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/shirakiyo/ConveniGacha/internal/domain/repository"

	"github.com/shirakiyo/ConveniGacha/internal/domain/model"
)

type productsCSV struct {
	filePath string
}

func NewProductRepository(path string) repository.ProductRepository {
	return &productsCSV{
		filePath: path,
	}
}

func (pr *productsCSV) ListProducts(fileName string) (result []*model.Product, err error) {
	result = make([]*model.Product, 0)

	csvFile, err := os.Open(filepath.Join(pr.filePath, fileName))
	if err != nil {
		return nil, err
	}
	defer func() {
		e := csvFile.Close()
		if e == nil {
			return
		}
		err = fmt.Errorf("Failed to close: %v, the original error was %v ", e, err)
	}()

	reader := csv.NewReader(csvFile)
	var line []string

	for i := 1; ; i++ {
		line, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if len(line) != 4 {
			return nil, fmt.Errorf("CSVファイルの%d行目に誤りがあります", i)
		}

		price, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, fmt.Errorf("CSVファイルの%d行目に誤りがあります: %w", i, err)
		}

		content := &model.Product{
			Name:   line[0],
			Price:  price,
			Link:   line[2],
			Detail: line[3],
		}

		result = append(result, content)
	}

	return result, nil
}
