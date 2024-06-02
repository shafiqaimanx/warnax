package show

import (
	"fmt"

	"github.com/shafiqaimanx/warnax/colors"
)

func StandardIndicators() {
	fmt.Printf("\t%s \t\t\tTEXT_ERROR\n", colors.TEXT_ERROR)
	fmt.Printf("\t%s \t\tTEXT_WARN\n", colors.TEXT_WARN)
	fmt.Printf("\t%s \t\t\tTEXT_OK\n", colors.TEXT_OK)
	fmt.Printf("\t%s \t\t\tTEXT_INFO\n\n", colors.TEXT_INFO)

	fmt.Printf("\t%s \t\t\tTEXT_SHORT_ERROR\n", colors.TEXT_SHORT_ERROR)
	fmt.Printf("\t%s \t\t\tTEXT_SHORT_WARN\n", colors.TEXT_SHORT_WARN)
	fmt.Printf("\t%s \t\t\tTEXT_SHORT_OK\n", colors.TEXT_SHORT_OK)
	fmt.Printf("\t%s \t\t\tTEXT_SHORT_INFO\n\n", colors.TEXT_SHORT_INFO)

	fmt.Printf("\t%s \t\t\tSQUARE_TEXT_SHORT_ERROR\n", colors.SQUARE_TEXT_SHORT_ERROR)
	fmt.Printf("\t%s \t\t\tSQUARE_TEXT_SHORT_WARN\n", colors.SQUARE_TEXT_SHORT_WARN)
	fmt.Printf("\t%s \t\t\tSQUARE_TEXT_SHORT_OK\n", colors.SQUARE_TEXT_SHORT_OK)
	fmt.Printf("\t%s \t\t\tSQUARE_TEXT_SHORT_INFO\n\n", colors.SQUARE_TEXT_SHORT_INFO)

	fmt.Printf("\t%s \t\t\tSQUARE_ERROR\n", colors.SQUARE_ERROR)
	fmt.Printf("\t%s \t\t\tSQUARE_WARN\n", colors.SQUARE_WARN)
	fmt.Printf("\t%s \t\t\tSQUARE_OK\n", colors.SQUARE_OK)
	fmt.Printf("\t%s \t\t\tSQUARE_INFO\n\n", colors.SQUARE_INFO)
}

func BoldIndicators() {
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_ERROR\n", colors.BOLD_TEXT_ERROR)
	fmt.Printf("\t%s \t\tBOLD_TEXT_WARN\n", colors.BOLD_TEXT_WARN)
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_OK\n", colors.BOLD_TEXT_OK)
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_INFO\n\n", colors.BOLD_TEXT_INFO)

	fmt.Printf("\t%s \t\t\tBOLD_TEXT_SHORT_ERROR\n", colors.BOLD_TEXT_SHORT_ERROR)
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_SHORT_WARN\n", colors.BOLD_TEXT_SHORT_WARN)
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_SHORT_OK\n", colors.BOLD_TEXT_SHORT_OK)
	fmt.Printf("\t%s \t\t\tBOLD_TEXT_SHORT_INFO\n\n", colors.BOLD_TEXT_SHORT_INFO)

	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_TEXT_SHORT_ERROR\n", colors.BOLD_SQUARE_TEXT_SHORT_ERROR)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_TEXT_SHORT_WARN\n", colors.BOLD_SQUARE_TEXT_SHORT_WARN)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_TEXT_SHORT_OK\n", colors.BOLD_SQUARE_TEXT_SHORT_OK)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_TEXT_SHORT_INFO\n\n", colors.BOLD_SQUARE_TEXT_SHORT_INFO)

	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_ERROR\n", colors.BOLD_SQUARE_ERROR)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_WARN\n", colors.BOLD_SQUARE_WARN)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_OK\n", colors.BOLD_SQUARE_OK)
	fmt.Printf("\t%s \t\t\tBOLD_SQUARE_INFO\n\n", colors.BOLD_SQUARE_INFO)
}