package generator

import (
	"github.com/jung-kurt/gofpdf"
	"net/url"
	"reflect"
	"testing"
)

func TestNewPDFGenerator(t *testing.T) {
	type args struct {
		data                MetaData
		strictErrorHandling bool
	}
	tests := []struct {
		name    string
		args    args
		wantGen *PDFGenerator
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGen, err := NewPDFGenerator(tt.args.data, tt.args.strictErrorHandling)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPDFGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGen, tt.wantGen) {
				t.Errorf("NewPDFGenerator() gotGen = %v, want %v", gotGen, tt.wantGen)
			}
		})
	}
}

func TestPDFGenerator_DrawLine(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		x1       float64
		y1       float64
		x2       float64
		y2       float64
		color    Color
		lineWith float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.DrawLine(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color, tt.args.lineWith)
		})
	}
}

func TestPDFGenerator_NewLine(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		oldX float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.NewLine(tt.args.oldX)
		})
	}
}

func TestPDFGenerator_PlaceMimeImageFromUrl(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		cdnUrl   *url.URL
		scale    float64
		alignStr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PlaceMimeImageFromUrl(tt.args.cdnUrl, tt.args.scale, tt.args.alignStr)
		})
	}
}

func TestPDFGenerator_PrintLnPdfText(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		text     string
		styleStr string
		alignStr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintLnPdfText(tt.args.text, tt.args.styleStr, tt.args.alignStr)
		})
	}
}

func TestPDFGenerator_PrintPdfText(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		text     string
		styleStr string
		alignStr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintPdfText(tt.args.text, tt.args.styleStr, tt.args.alignStr)
		})
	}
}

func TestPDFGenerator_PrintPdfTextFormatted(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		text            string
		styleStr        string
		alignStr        string
		borderStr       string
		fill            bool
		backgroundColor Color
		cellHeight      float64
		cellWidth       float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintPdfTextFormatted(tt.args.text, tt.args.styleStr, tt.args.alignStr, tt.args.borderStr, tt.args.fill, tt.args.backgroundColor, tt.args.cellHeight, tt.args.cellWidth)
		})
	}
}

func TestPDFGenerator_PrintTableBody(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		cells              [][]string
		columnWidths       []float64
		columnAlignStrings []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintTableBody(tt.args.cells, tt.args.columnWidths, tt.args.columnAlignStrings)
		})
	}
}

func TestPDFGenerator_PrintTableFooter(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		cells              [][]string
		columnWidths       []float64
		columnAlignStrings []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintTableFooter(tt.args.cells, tt.args.columnWidths, tt.args.columnAlignStrings)
		})
	}
}

func TestPDFGenerator_PrintTableHeader(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		cells              []string
		columnWidth        []float64
		columnAlignStrings []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.PrintTableHeader(tt.args.cells, tt.args.columnWidth, tt.args.columnAlignStrings)
		})
	}
}

func TestPDFGenerator_addNewPageIfNecessary(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.addNewPageIfNecessary()
		})
	}
}

func TestPDFGenerator_extractLinesFromText(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		text string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantTextLines []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			if gotTextLines := core.extractLinesFromText(tt.args.text); !reflect.DeepEqual(gotTextLines, tt.wantTextLines) {
				t.Errorf("extractLinesFromText() = %v, want %v", gotTextLines, tt.wantTextLines)
			}
		})
	}
}

func TestPDFGenerator_printTableBodyRow(t *testing.T) {
	type fields struct {
		pdf                 *gofpdf.Fpdf
		data                MetaData
		maxSaveX            float64
		maxSaveY            float64
		strictErrorHandling bool
	}
	type args struct {
		extractedLines [][]string
		currentLine    int
		maxItems       int
		alignStrings   []string
		newlineHeight  float64
		columnWidth    []float64
		referenceX     float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := &PDFGenerator{
				pdf:                 tt.fields.pdf,
				data:                tt.fields.data,
				maxSaveX:            tt.fields.maxSaveX,
				maxSaveY:            tt.fields.maxSaveY,
				strictErrorHandling: tt.fields.strictErrorHandling,
			}
			core.printTableBodyRow(tt.args.extractedLines, tt.args.currentLine, tt.args.maxItems, tt.args.alignStrings, tt.args.newlineHeight, tt.args.columnWidth, tt.args.referenceX)
		})
	}
}