package freemono

import "tinygo.org/x/tinyfont"

var Regular24pt7bBitmaps = []byte{
	0x73, 0x9C, 0xE7, 0x39, 0xCE, 0x73, 0x9C, 0xE7, 0x10, 0x84, 0x21, 0x08,
	0x00, 0x00, 0x00, 0x03, 0xBF, 0xFF, 0xB8, 0xFE, 0x7F, 0x7C, 0x3E, 0x7C,
	0x3E, 0x7C, 0x3E, 0x7C, 0x3E, 0x7C, 0x3E, 0x7C, 0x3E, 0x7C, 0x3E, 0x3C,
	0x3E, 0x38, 0x1C, 0x38, 0x1C, 0x38, 0x1C, 0x38, 0x1C, 0x38, 0x1C, 0x01,
	0x86, 0x00, 0x30, 0xC0, 0x06, 0x18, 0x00, 0xC3, 0x00, 0x18, 0x60, 0x03,
	0x0C, 0x00, 0x61, 0x80, 0x0C, 0x70, 0x01, 0x8C, 0x00, 0x61, 0x80, 0x0C,
	0x30, 0x3F, 0xFF, 0xF7, 0xFF, 0xFE, 0x06, 0x18, 0x00, 0xC3, 0x00, 0x18,
	0x60, 0x03, 0x0C, 0x00, 0x61, 0x80, 0x0C, 0x30, 0x7F, 0xFF, 0xEF, 0xFF,
	0xFC, 0x06, 0x18, 0x00, 0xC7, 0x00, 0x38, 0xC0, 0x06, 0x18, 0x00, 0xC3,
	0x00, 0x18, 0x60, 0x03, 0x0C, 0x00, 0x61, 0x80, 0x0C, 0x30, 0x01, 0x86,
	0x00, 0x30, 0xC0, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x0F, 0xC0,
	0x0F, 0xFD, 0x87, 0x03, 0xE3, 0x80, 0x39, 0xC0, 0x06, 0x60, 0x01, 0x98,
	0x00, 0x06, 0x00, 0x01, 0xC0, 0x00, 0x38, 0x00, 0x07, 0xC0, 0x00, 0x7F,
	0x80, 0x03, 0xF8, 0x00, 0x0F, 0x80, 0x00, 0x60, 0x00, 0x1C, 0x00, 0x03,
	0x80, 0x00, 0xF0, 0x00, 0x3C, 0x00, 0x1F, 0x80, 0x0E, 0xFC, 0x0F, 0x37,
	0xFF, 0x80, 0x7F, 0x80, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C,
	0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x07, 0x80, 0x01, 0xFE, 0x00, 0x38,
	0x70, 0x03, 0x03, 0x00, 0x60, 0x18, 0x06, 0x01, 0x80, 0x60, 0x18, 0x06,
	0x01, 0x80, 0x30, 0x30, 0x03, 0x87, 0x00, 0x1F, 0xE0, 0x30, 0x78, 0x1F,
	0x00, 0x1F, 0x80, 0x0F, 0xC0, 0x07, 0xE0, 0x03, 0xF0, 0x00, 0xF8, 0x00,
	0x0C, 0x01, 0xE0, 0x00, 0x7F, 0x80, 0x0E, 0x1C, 0x00, 0xC0, 0xC0, 0x18,
	0x06, 0x01, 0x80, 0x60, 0x18, 0x06, 0x01, 0x80, 0x60, 0x0C, 0x0E, 0x00,
	0xE1, 0xC0, 0x07, 0xF8, 0x00, 0x1E, 0x00, 0x03, 0xEC, 0x01, 0xFF, 0x00,
	0xE1, 0x00, 0x70, 0x00, 0x18, 0x00, 0x06, 0x00, 0x01, 0x80, 0x00, 0x30,
	0x00, 0x0C, 0x00, 0x01, 0x80, 0x00, 0x60, 0x00, 0x7C, 0x00, 0x3B, 0x83,
	0xD8, 0x60, 0xFE, 0x0C, 0x33, 0x03, 0x98, 0xC0, 0x66, 0x30, 0x0D, 0x8C,
	0x03, 0xC3, 0x00, 0x70, 0x60, 0x1C, 0x1C, 0x0F, 0x03, 0x87, 0x7C, 0x7F,
	0x9F, 0x07, 0x80, 0x00, 0xFE, 0xF9, 0xF3, 0xE7, 0xCF, 0x9F, 0x3E, 0x3C,
	0x70, 0xE1, 0xC3, 0x87, 0x00, 0x06, 0x1C, 0x30, 0xE1, 0x87, 0x0E, 0x18,
	0x70, 0xE1, 0xC3, 0x0E, 0x1C, 0x38, 0x70, 0xE1, 0xC3, 0x87, 0x0E, 0x0C,
	0x1C, 0x38, 0x70, 0x60, 0xE1, 0xC1, 0x83, 0x83, 0x06, 0x06, 0x04, 0xC1,
	0xC1, 0x83, 0x83, 0x07, 0x0E, 0x0C, 0x1C, 0x38, 0x70, 0xE0, 0xE1, 0xC3,
	0x87, 0x0E, 0x1C, 0x38, 0x70, 0xE1, 0x87, 0x0E, 0x1C, 0x30, 0x61, 0xC3,
	0x0E, 0x18, 0x70, 0xC1, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x00,
	0x03, 0x00, 0x00, 0xC0, 0x10, 0x30, 0x3F, 0x8C, 0x7C, 0xFF, 0xFC, 0x07,
	0xF8, 0x00, 0x78, 0x00, 0x1F, 0x00, 0x0C, 0xC0, 0x06, 0x18, 0x03, 0x87,
	0x00, 0xC0, 0xC0, 0x60, 0x18, 0x00, 0x60, 0x00, 0x06, 0x00, 0x00, 0x60,
	0x00, 0x06, 0x00, 0x00, 0x60, 0x00, 0x06, 0x00, 0x00, 0x60, 0x00, 0x06,
	0x00, 0x00, 0x60, 0x00, 0x06, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00,
	0x60, 0x00, 0x06, 0x00, 0x00, 0x60, 0x00, 0x06, 0x00, 0x00, 0x60, 0x00,
	0x06, 0x00, 0x00, 0x60, 0x00, 0x06, 0x00, 0x00, 0x60, 0x00, 0x06, 0x00,
	0x1F, 0x8F, 0x87, 0xC7, 0xC3, 0xE1, 0xE1, 0xF0, 0xF0, 0x78, 0x38, 0x3C,
	0x1C, 0x0E, 0x06, 0x00, 0x7F, 0xFF, 0xFD, 0xFF, 0xFF, 0xF0, 0x7D, 0xFF,
	0xFF, 0xFF, 0xEF, 0x80, 0x00, 0x00, 0xC0, 0x00, 0x70, 0x00, 0x18, 0x00,
	0x06, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x60, 0x00, 0x18, 0x00, 0x0C,
	0x00, 0x03, 0x00, 0x01, 0x80, 0x00, 0x60, 0x00, 0x30, 0x00, 0x0C, 0x00,
	0x06, 0x00, 0x01, 0x80, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x18, 0x00, 0x06,
	0x00, 0x03, 0x80, 0x00, 0xC0, 0x00, 0x70, 0x00, 0x18, 0x00, 0x0E, 0x00,
	0x03, 0x00, 0x01, 0xC0, 0x00, 0x60, 0x00, 0x38, 0x00, 0x0C, 0x00, 0x07,
	0x00, 0x01, 0x80, 0x00, 0x60, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x00, 0x03,
	0xF0, 0x03, 0xFF, 0x01, 0xE1, 0xE0, 0xE0, 0x18, 0x30, 0x03, 0x1C, 0x00,
	0xE6, 0x00, 0x19, 0x80, 0x06, 0xE0, 0x01, 0xF0, 0x00, 0x3C, 0x00, 0x0F,
	0x00, 0x03, 0xC0, 0x00, 0xF0, 0x00, 0x3C, 0x00, 0x0F, 0x00, 0x03, 0xC0,
	0x00, 0xF0, 0x00, 0x3C, 0x00, 0x0F, 0x00, 0x03, 0xC0, 0x00, 0xF8, 0x00,
	0x76, 0x00, 0x19, 0x80, 0x06, 0x70, 0x03, 0x8C, 0x00, 0xC3, 0x80, 0x60,
	0x78, 0x78, 0x0F, 0xFC, 0x00, 0xFC, 0x00, 0x03, 0x80, 0x07, 0x80, 0x0F,
	0x80, 0x1D, 0x80, 0x39, 0x80, 0x71, 0x80, 0xE1, 0x80, 0xC1, 0x80, 0x01,
	0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01,
	0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01,
	0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01,
	0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0x03, 0xF0, 0x03, 0xFF, 0x01, 0xC0, 0xE0,
	0xC0, 0x1C, 0x60, 0x03, 0xB8, 0x00, 0x6C, 0x00, 0x0F, 0x00, 0x03, 0x00,
	0x00, 0xC0, 0x00, 0x30, 0x00, 0x18, 0x00, 0x06, 0x00, 0x03, 0x00, 0x01,
	0x80, 0x00, 0xC0, 0x00, 0x60, 0x00, 0x30, 0x00, 0x18, 0x00, 0x0C, 0x00,
	0x06, 0x00, 0x03, 0x00, 0x01, 0x80, 0x00, 0xC0, 0x00, 0x60, 0x00, 0x30,
	0x00, 0xD0, 0x00, 0x38, 0x00, 0x0F, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0x03,
	0xF8, 0x01, 0xFF, 0xC0, 0x70, 0x3C, 0x18, 0x01, 0xC6, 0x00, 0x18, 0x00,
	0x01, 0x80, 0x00, 0x30, 0x00, 0x06, 0x00, 0x00, 0xC0, 0x00, 0x18, 0x00,
	0x06, 0x00, 0x01, 0xC0, 0x00, 0x70, 0x01, 0xFC, 0x00, 0x3F, 0x00, 0x00,
	0x78, 0x00, 0x03, 0x80, 0x00, 0x38, 0x00, 0x03, 0x00, 0x00, 0x30, 0x00,
	0x06, 0x00, 0x00, 0xC0, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00, 0xD8, 0x00,
	0x3B, 0x80, 0x0E, 0x3E, 0x07, 0x81, 0xFF, 0xE0, 0x07, 0xE0, 0x00, 0x00,
	0x3C, 0x00, 0x7C, 0x00, 0x6C, 0x00, 0xCC, 0x00, 0x8C, 0x01, 0x8C, 0x03,
	0x0C, 0x03, 0x0C, 0x06, 0x0C, 0x04, 0x0C, 0x0C, 0x0C, 0x08, 0x0C, 0x10,
	0x0C, 0x30, 0x0C, 0x20, 0x0C, 0x60, 0x0C, 0x40, 0x0C, 0x80, 0x0C, 0xFF,
	0xFF, 0xFF, 0xFF, 0x00, 0x0C, 0x00, 0x0C, 0x00, 0x0C, 0x00, 0x0C, 0x00,
	0x0C, 0x00, 0x0C, 0x00, 0xFF, 0x00, 0xFF, 0x3F, 0xFF, 0x07, 0xFF, 0xE0,
	0xC0, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x0C, 0x00, 0x01,
	0x80, 0x00, 0x30, 0x00, 0x06, 0x00, 0x00, 0xC7, 0xE0, 0x1F, 0xFF, 0x03,
	0x80, 0x70, 0x00, 0x03, 0x00, 0x00, 0x30, 0x00, 0x06, 0x00, 0x00, 0x60,
	0x00, 0x0C, 0x00, 0x01, 0x80, 0x00, 0x30, 0x00, 0x06, 0x00, 0x00, 0xC0,
	0x00, 0x30, 0x00, 0x06, 0xC0, 0x01, 0xDC, 0x00, 0x71, 0xF0, 0x3C, 0x0F,
	0xFF, 0x00, 0x3F, 0x00, 0x00, 0x3F, 0x80, 0x3F, 0xF0, 0x3E, 0x00, 0x1E,
	0x00, 0x0E, 0x00, 0x07, 0x00, 0x03, 0x80, 0x00, 0xC0, 0x00, 0x70, 0x00,
	0x18, 0x00, 0x06, 0x00, 0x03, 0x80, 0x00, 0xC1, 0xF8, 0x31, 0xFF, 0x0C,
	0xF0, 0xF3, 0x70, 0x0C, 0xD8, 0x01, 0xBC, 0x00, 0x6E, 0x00, 0x0F, 0x80,
	0x03, 0xC0, 0x00, 0xD8, 0x00, 0x36, 0x00, 0x0D, 0x80, 0x03, 0x30, 0x01,
	0x8E, 0x00, 0x61, 0xC0, 0x30, 0x38, 0x38, 0x07, 0xFC, 0x00, 0x7C, 0x00,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x00, 0x0F, 0x00, 0x03, 0xC0, 0x01, 0xC0,
	0x00, 0x60, 0x00, 0x18, 0x00, 0x0E, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00,
	0x30, 0x00, 0x18, 0x00, 0x06, 0x00, 0x01, 0x80, 0x00, 0xC0, 0x00, 0x30,
	0x00, 0x0C, 0x00, 0x06, 0x00, 0x01, 0x80, 0x00, 0x60, 0x00, 0x30, 0x00,
	0x0C, 0x00, 0x03, 0x00, 0x01, 0x80, 0x00, 0x60, 0x00, 0x18, 0x00, 0x0C,
	0x00, 0x03, 0x00, 0x03, 0xF0, 0x03, 0xFF, 0x03, 0xC0, 0xF1, 0xC0, 0x0E,
	0x60, 0x01, 0xB8, 0x00, 0x7C, 0x00, 0x0F, 0x00, 0x03, 0xC0, 0x00, 0xF0,
	0x00, 0x36, 0x00, 0x18, 0xC0, 0x0C, 0x1C, 0x0E, 0x03, 0xFF, 0x00, 0xFF,
	0xC0, 0x70, 0x38, 0x30, 0x03, 0x18, 0x00, 0x66, 0x00, 0x1B, 0x00, 0x03,
	0xC0, 0x00, 0xF0, 0x00, 0x3C, 0x00, 0x0F, 0x00, 0x03, 0x60, 0x01, 0x98,
	0x00, 0xE3, 0x00, 0x70, 0x70, 0x38, 0x0F, 0xFC, 0x00, 0xFC, 0x00, 0x07,
	0xE0, 0x03, 0xFE, 0x01, 0xC1, 0xC0, 0xC0, 0x38, 0x60, 0x07, 0x18, 0x00,
	0xCC, 0x00, 0x1B, 0x00, 0x06, 0xC0, 0x01, 0xB0, 0x00, 0x3C, 0x00, 0x1F,
	0x00, 0x07, 0x60, 0x03, 0xD8, 0x01, 0xB3, 0x00, 0xCC, 0xF0, 0xF3, 0x0F,
	0xF8, 0xC1, 0xF8, 0x30, 0x00, 0x1C, 0x00, 0x06, 0x00, 0x01, 0x80, 0x00,
	0xE0, 0x00, 0x30, 0x00, 0x1C, 0x00, 0x0E, 0x00, 0x07, 0x00, 0x07, 0x80,
	0x07, 0xC0, 0xFF, 0xC0, 0x1F, 0xC0, 0x00, 0x7D, 0xFF, 0xFF, 0xFF, 0xEF,
	0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0xFF, 0xFF, 0xFF,
	0xF7, 0xC0, 0x0F, 0x87, 0xF1, 0xFC, 0x7F, 0x1F, 0xC3, 0xE0, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0xF1, 0xF8, 0x7C, 0x3F, 0x0F,
	0x83, 0xE0, 0xF0, 0x7C, 0x1E, 0x07, 0x81, 0xC0, 0xF0, 0x38, 0x04, 0x00,
	0x00, 0x00, 0x18, 0x00, 0x01, 0xE0, 0x00, 0x1E, 0x00, 0x00, 0xE0, 0x00,
	0x0F, 0x00, 0x00, 0xF0, 0x00, 0x0F, 0x00, 0x00, 0xF0, 0x00, 0x07, 0x00,
	0x00, 0x78, 0x00, 0x07, 0x80, 0x00, 0x0F, 0x00, 0x00, 0x1E, 0x00, 0x00,
	0x1E, 0x00, 0x00, 0x1E, 0x00, 0x00, 0x1E, 0x00, 0x00, 0x1E, 0x00, 0x00,
	0x3C, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x38, 0x00, 0x00,
	0x20, 0x7F, 0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xFF,
	0xFF, 0x7F, 0xFF, 0xFF, 0xC0, 0x00, 0x07, 0x80, 0x00, 0x0F, 0x00, 0x00,
	0x1E, 0x00, 0x00, 0x38, 0x00, 0x00, 0xF0, 0x00, 0x01, 0xE0, 0x00, 0x03,
	0xC0, 0x00, 0x07, 0x80, 0x00, 0x0E, 0x00, 0x00, 0x3C, 0x00, 0x01, 0xE0,
	0x00, 0x3C, 0x00, 0x07, 0x80, 0x00, 0xF0, 0x00, 0x1E, 0x00, 0x01, 0xE0,
	0x00, 0x3C, 0x00, 0x07, 0x80, 0x00, 0xF0, 0x00, 0x0E, 0x00, 0x00, 0x60,
	0x00, 0x00, 0x07, 0xF0, 0x1F, 0xFE, 0x3E, 0x07, 0x98, 0x00, 0xEC, 0x00,
	0x36, 0x00, 0x0F, 0x00, 0x06, 0x00, 0x03, 0x00, 0x01, 0x80, 0x01, 0xC0,
	0x00, 0xC0, 0x01, 0xC0, 0x03, 0xC0, 0x07, 0xC0, 0x07, 0x00, 0x03, 0x00,
	0x01, 0x80, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x07, 0x80, 0x07, 0xE0, 0x03, 0xF0, 0x01, 0xF8, 0x00,
	0x78, 0x00, 0x03, 0xF0, 0x03, 0xFF, 0x01, 0xE0, 0xE0, 0xE0, 0x1C, 0x30,
	0x03, 0x1C, 0x00, 0x66, 0x00, 0x19, 0x80, 0x06, 0xC0, 0x01, 0xB0, 0x07,
	0xEC, 0x07, 0xFB, 0x03, 0xC6, 0xC1, 0xC1, 0xB0, 0xE0, 0x6C, 0x30, 0x1B,
	0x0C, 0x06, 0xC3, 0x01, 0xB0, 0xC0, 0x6C, 0x18, 0x1B, 0x07, 0x86, 0xC0,
	0xFF, 0xF0, 0x0F, 0xFC, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x18, 0x00,
	0x07, 0x00, 0x00, 0xC0, 0x00, 0x38, 0x00, 0x07, 0x80, 0xC0, 0xFF, 0xF0,
	0x0F, 0xE0, 0x07, 0xFF, 0x00, 0x00, 0x7F, 0xF0, 0x00, 0x00, 0x1B, 0x00,
	0x00, 0x01, 0x98, 0x00, 0x00, 0x11, 0x80, 0x00, 0x03, 0x0C, 0x00, 0x00,
	0x30, 0xC0, 0x00, 0x06, 0x0C, 0x00, 0x00, 0x60, 0x60, 0x00, 0x06, 0x06,
	0x00, 0x00, 0xC0, 0x30, 0x00, 0x0C, 0x03, 0x00, 0x00, 0x80, 0x30, 0x00,
	0x18, 0x01, 0x80, 0x01, 0x80, 0x18, 0x00, 0x3F, 0xFF, 0x80, 0x03, 0xFF,
	0xFC, 0x00, 0x20, 0x00, 0xC0, 0x06, 0x00, 0x06, 0x00, 0x60, 0x00, 0x60,
	0x0C, 0x00, 0x06, 0x00, 0xC0, 0x00, 0x30, 0x0C, 0x00, 0x03, 0x01, 0x80,
	0x00, 0x18, 0x7F, 0xC0, 0x3F, 0xF7, 0xFC, 0x03, 0xFF, 0xFF, 0xFF, 0x03,
	0xFF, 0xFF, 0x01, 0x80, 0x0E, 0x06, 0x00, 0x1C, 0x18, 0x00, 0x38, 0x60,
	0x00, 0x61, 0x80, 0x01, 0x86, 0x00, 0x06, 0x18, 0x00, 0x38, 0x60, 0x01,
	0xC1, 0x80, 0x1E, 0x07, 0xFF, 0xE0, 0x1F, 0xFF, 0xC0, 0x60, 0x03, 0xC1,
	0x80, 0x03, 0x86, 0x00, 0x06, 0x18, 0x00, 0x1C, 0x60, 0x00, 0x31, 0x80,
	0x00, 0xC6, 0x00, 0x03, 0x18, 0x00, 0x0C, 0x60, 0x00, 0x61, 0x80, 0x03,
	0x86, 0x00, 0x1C, 0xFF, 0xFF, 0xE3, 0xFF, 0xFE, 0x00, 0x00, 0xFC, 0x00,
	0x0F, 0xFE, 0x60, 0xF0, 0x3D, 0x87, 0x00, 0x3E, 0x38, 0x00, 0x38, 0xC0,
	0x00, 0xE7, 0x00, 0x01, 0x98, 0x00, 0x06, 0x60, 0x00, 0x03, 0x00, 0x00,
	0x0C, 0x00, 0x00, 0x30, 0x00, 0x00, 0xC0, 0x00, 0x03, 0x00, 0x00, 0x0C,
	0x00, 0x00, 0x30, 0x00, 0x00, 0xC0, 0x00, 0x03, 0x00, 0x00, 0x0C, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x60, 0x00, 0x01, 0xC0, 0x00, 0x03, 0x80, 0x00,
	0xC7, 0x00, 0x06, 0x0E, 0x00, 0x70, 0x1E, 0x07, 0x80, 0x3F, 0xFC, 0x00,
	0x1F, 0x80, 0xFF, 0xFE, 0x03, 0xFF, 0xFE, 0x03, 0x00, 0x3C, 0x0C, 0x00,
	0x38, 0x30, 0x00, 0x70, 0xC0, 0x00, 0xC3, 0x00, 0x03, 0x8C, 0x00, 0x06,
	0x30, 0x00, 0x1C, 0xC0, 0x00, 0x33, 0x00, 0x00, 0xCC, 0x00, 0x03, 0x30,
	0x00, 0x0C, 0xC0, 0x00, 0x33, 0x00, 0x00, 0xCC, 0x00, 0x03, 0x30, 0x00,
	0x0C, 0xC0, 0x00, 0x33, 0x00, 0x01, 0x8C, 0x00, 0x06, 0x30, 0x00, 0x30,
	0xC0, 0x01, 0xC3, 0x00, 0x0E, 0x0C, 0x00, 0xF0, 0xFF, 0xFF, 0x83, 0xFF,
	0xF8, 0x00, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xE1, 0x80, 0x01, 0x86, 0x00,
	0x06, 0x18, 0x00, 0x18, 0x60, 0x00, 0x61, 0x80, 0x01, 0x86, 0x00, 0x00,
	0x18, 0x0C, 0x00, 0x60, 0x30, 0x01, 0x80, 0xC0, 0x07, 0xFF, 0x00, 0x1F,
	0xFC, 0x00, 0x60, 0x30, 0x01, 0x80, 0xC0, 0x06, 0x03, 0x00, 0x18, 0x00,
	0x00, 0x60, 0x00, 0x01, 0x80, 0x00, 0xC6, 0x00, 0x03, 0x18, 0x00, 0x0C,
	0x60, 0x00, 0x31, 0x80, 0x00, 0xC6, 0x00, 0x03, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xF0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF1, 0x80, 0x00, 0xC6, 0x00,
	0x03, 0x18, 0x00, 0x0C, 0x60, 0x00, 0x31, 0x80, 0x00, 0xC6, 0x00, 0x00,
	0x18, 0x0C, 0x00, 0x60, 0x30, 0x01, 0x80, 0xC0, 0x07, 0xFF, 0x00, 0x1F,
	0xFC, 0x00, 0x60, 0x30, 0x01, 0x80, 0xC0, 0x06, 0x03, 0x00, 0x18, 0x00,
	0x00, 0x60, 0x00, 0x01, 0x80, 0x00, 0x06, 0x00, 0x00, 0x18, 0x00, 0x00,
	0x60, 0x00, 0x01, 0x80, 0x00, 0x06, 0x00, 0x00, 0xFF, 0xF0, 0x03, 0xFF,
	0xC0, 0x00, 0x00, 0xFF, 0x00, 0x07, 0xFF, 0x98, 0x1E, 0x03, 0xF0, 0x70,
	0x01, 0xE1, 0x80, 0x01, 0xC6, 0x00, 0x01, 0x9C, 0x00, 0x03, 0x30, 0x00,
	0x00, 0x60, 0x00, 0x01, 0xC0, 0x00, 0x03, 0x00, 0x00, 0x06, 0x00, 0x00,
	0x0C, 0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0x60, 0x03, 0xFF,
	0xC0, 0x07, 0xFF, 0x80, 0x00, 0x1B, 0x00, 0x00, 0x37, 0x00, 0x00, 0x66,
	0x00, 0x00, 0xCC, 0x00, 0x01, 0x8C, 0x00, 0x03, 0x1C, 0x00, 0x06, 0x1E,
	0x00, 0x0C, 0x0F, 0x00, 0xF8, 0x0F, 0xFF, 0xC0, 0x03, 0xFC, 0x00, 0x7F,
	0x01, 0xFC, 0xFE, 0x03, 0xF8, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80,
	0x03, 0x03, 0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00,
	0x30, 0x30, 0x00, 0x60, 0x60, 0x00, 0xC0, 0xFF, 0xFF, 0x81, 0xFF, 0xFF,
	0x03, 0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00, 0x30,
	0x30, 0x00, 0x60, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80, 0x03, 0x03,
	0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0xFF, 0x01, 0xFF, 0xFE,
	0x03, 0xFC, 0xFF, 0xFF, 0xFF, 0xFF, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80,
	0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80,
	0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80,
	0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80,
	0x01, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFE, 0x01, 0xFF, 0xFC,
	0x00, 0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x18, 0x00,
	0x00, 0x30, 0x00, 0x00, 0x60, 0x00, 0x00, 0xC0, 0x00, 0x01, 0x80, 0x00,
	0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x18, 0x00, 0x00,
	0x30, 0x60, 0x00, 0x60, 0xC0, 0x00, 0xC1, 0x80, 0x01, 0x83, 0x00, 0x03,
	0x06, 0x00, 0x06, 0x0C, 0x00, 0x0C, 0x18, 0x00, 0x30, 0x38, 0x00, 0x60,
	0x38, 0x01, 0x80, 0x3C, 0x0E, 0x00, 0x3F, 0xF8, 0x00, 0x0F, 0xC0, 0x00,
	0xFF, 0x81, 0xFE, 0xFF, 0x81, 0xFE, 0x18, 0x00, 0x30, 0x18, 0x00, 0xE0,
	0x18, 0x01, 0xC0, 0x18, 0x03, 0x80, 0x18, 0x07, 0x00, 0x18, 0x0E, 0x00,
	0x18, 0x18, 0x00, 0x18, 0x70, 0x00, 0x18, 0xE0, 0x00, 0x19, 0xE0, 0x00,
	0x1B, 0xF8, 0x00, 0x1F, 0x1C, 0x00, 0x1C, 0x06, 0x00, 0x18, 0x03, 0x00,
	0x18, 0x03, 0x80, 0x18, 0x01, 0x80, 0x18, 0x00, 0xC0, 0x18, 0x00, 0xC0,
	0x18, 0x00, 0x60, 0x18, 0x00, 0x60, 0x18, 0x00, 0x70, 0x18, 0x00, 0x30,
	0xFF, 0x80, 0x3F, 0xFF, 0x80, 0x1F, 0xFF, 0xF0, 0x07, 0xFF, 0x80, 0x01,
	0x80, 0x00, 0x0C, 0x00, 0x00, 0x60, 0x00, 0x03, 0x00, 0x00, 0x18, 0x00,
	0x00, 0xC0, 0x00, 0x06, 0x00, 0x00, 0x30, 0x00, 0x01, 0x80, 0x00, 0x0C,
	0x00, 0x00, 0x60, 0x00, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0xC0, 0x00,
	0x06, 0x00, 0x18, 0x30, 0x00, 0xC1, 0x80, 0x06, 0x0C, 0x00, 0x30, 0x60,
	0x01, 0x83, 0x00, 0x0C, 0x18, 0x00, 0x60, 0xC0, 0x03, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xC0, 0xFC, 0x00, 0x0F, 0xFF, 0x00, 0x03, 0xF3, 0x60, 0x01,
	0xB0, 0xD8, 0x00, 0x6C, 0x33, 0x00, 0x33, 0x0C, 0xC0, 0x0C, 0xC3, 0x38,
	0x07, 0x30, 0xC6, 0x01, 0x8C, 0x31, 0xC0, 0xE3, 0x0C, 0x30, 0x30, 0xC3,
	0x0C, 0x0C, 0x30, 0xC1, 0x86, 0x0C, 0x30, 0x61, 0x83, 0x0C, 0x0C, 0xC0,
	0xC3, 0x03, 0x30, 0x30, 0xC0, 0x78, 0x0C, 0x30, 0x1E, 0x03, 0x0C, 0x03,
	0x00, 0xC3, 0x00, 0x00, 0x30, 0xC0, 0x00, 0x0C, 0x30, 0x00, 0x03, 0x0C,
	0x00, 0x00, 0xC3, 0x00, 0x00, 0x30, 0xC0, 0x00, 0x0C, 0xFF, 0x00, 0x3F,
	0xFF, 0xC0, 0x0F, 0xF0, 0xFC, 0x00, 0xFF, 0xFC, 0x00, 0xFF, 0x1E, 0x00,
	0x0C, 0x1F, 0x00, 0x0C, 0x1B, 0x00, 0x0C, 0x19, 0x80, 0x0C, 0x19, 0xC0,
	0x0C, 0x18, 0xC0, 0x0C, 0x18, 0x60, 0x0C, 0x18, 0x60, 0x0C, 0x18, 0x30,
	0x0C, 0x18, 0x38, 0x0C, 0x18, 0x18, 0x0C, 0x18, 0x0C, 0x0C, 0x18, 0x0E,
	0x0C, 0x18, 0x06, 0x0C, 0x18, 0x03, 0x0C, 0x18, 0x03, 0x0C, 0x18, 0x01,
	0x8C, 0x18, 0x01, 0xCC, 0x18, 0x00, 0xCC, 0x18, 0x00, 0x6C, 0x18, 0x00,
	0x7C, 0x18, 0x00, 0x3C, 0x7F, 0x80, 0x1C, 0x7F, 0x80, 0x1C, 0x00, 0x7E,
	0x00, 0x01, 0xFF, 0xC0, 0x07, 0x81, 0xE0, 0x0E, 0x00, 0x70, 0x1C, 0x00,
	0x38, 0x38, 0x00, 0x1C, 0x30, 0x00, 0x0C, 0x70, 0x00, 0x0E, 0x60, 0x00,
	0x06, 0x60, 0x00, 0x06, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00,
	0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00,
	0x03, 0xC0, 0x00, 0x03, 0x60, 0x00, 0x06, 0x60, 0x00, 0x06, 0x70, 0x00,
	0x0E, 0x30, 0x00, 0x0C, 0x38, 0x00, 0x1C, 0x1C, 0x00, 0x38, 0x0E, 0x00,
	0x70, 0x07, 0x81, 0xE0, 0x03, 0xFF, 0xC0, 0x00, 0x7E, 0x00, 0xFF, 0xFF,
	0x07, 0xFF, 0xFE, 0x06, 0x00, 0x78, 0x30, 0x00, 0xE1, 0x80, 0x03, 0x0C,
	0x00, 0x0C, 0x60, 0x00, 0x63, 0x00, 0x03, 0x18, 0x00, 0x18, 0xC0, 0x01,
	0xC6, 0x00, 0x0C, 0x30, 0x00, 0xC1, 0x80, 0x1E, 0x0F, 0xFF, 0xC0, 0x7F,
	0xF8, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0xC0, 0x00, 0x06, 0x00, 0x00,
	0x30, 0x00, 0x01, 0x80, 0x00, 0x0C, 0x00, 0x00, 0x60, 0x00, 0x03, 0x00,
	0x00, 0xFF, 0xF0, 0x07, 0xFF, 0x80, 0x00, 0x00, 0x7E, 0x00, 0x01, 0xFF,
	0x80, 0x07, 0x81, 0xE0, 0x0E, 0x00, 0x70, 0x1C, 0x00, 0x38, 0x38, 0x00,
	0x1C, 0x30, 0x00, 0x0C, 0x70, 0x00, 0x0E, 0x60, 0x00, 0x06, 0x60, 0x00,
	0x06, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00,
	0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00, 0x03, 0xC0, 0x00,
	0x03, 0x60, 0x00, 0x06, 0x60, 0x00, 0x06, 0x70, 0x00, 0x0E, 0x30, 0x00,
	0x0C, 0x18, 0x00, 0x1C, 0x0C, 0x00, 0x38, 0x06, 0x00, 0x70, 0x03, 0x81,
	0xE0, 0x00, 0xFF, 0xC0, 0x00, 0x7E, 0x00, 0x00, 0xE0, 0x00, 0x03, 0xFF,
	0x87, 0x07, 0xFF, 0xFE, 0x07, 0x00, 0xF8, 0xFF, 0xFE, 0x00, 0xFF, 0xFF,
	0x80, 0x18, 0x03, 0xC0, 0x18, 0x00, 0xE0, 0x18, 0x00, 0x60, 0x18, 0x00,
	0x30, 0x18, 0x00, 0x30, 0x18, 0x00, 0x30, 0x18, 0x00, 0x30, 0x18, 0x00,
	0x70, 0x18, 0x00, 0x60, 0x18, 0x01, 0xC0, 0x18, 0x07, 0x80, 0x1F, 0xFF,
	0x00, 0x1F, 0xFC, 0x00, 0x18, 0x0E, 0x00, 0x18, 0x07, 0x00, 0x18, 0x03,
	0x80, 0x18, 0x01, 0xC0, 0x18, 0x00, 0xE0, 0x18, 0x00, 0x60, 0x18, 0x00,
	0x30, 0x18, 0x00, 0x30, 0x18, 0x00, 0x18, 0xFF, 0x80, 0x1F, 0xFF, 0x80,
	0x0F, 0x03, 0xF8, 0x00, 0xFF, 0xE6, 0x1E, 0x07, 0xE3, 0x80, 0x1E, 0x30,
	0x00, 0xE6, 0x00, 0x06, 0x60, 0x00, 0x66, 0x00, 0x06, 0x60, 0x00, 0x07,
	0x00, 0x00, 0x30, 0x00, 0x01, 0xC0, 0x00, 0x0F, 0xC0, 0x00, 0x3F, 0xC0,
	0x00, 0x3F, 0x80, 0x00, 0x1C, 0x00, 0x00, 0xE0, 0x00, 0x07, 0x00, 0x00,
	0x30, 0x00, 0x03, 0xC0, 0x00, 0x3C, 0x00, 0x03, 0xE0, 0x00, 0x7E, 0x00,
	0x06, 0xF8, 0x01, 0xED, 0xE0, 0x7C, 0xCF, 0xFF, 0x00, 0x3F, 0xC0, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFC, 0x03, 0x00, 0xF0, 0x0C, 0x03, 0xC0, 0x30,
	0x0F, 0x00, 0xC0, 0x3C, 0x03, 0x00, 0xC0, 0x0C, 0x00, 0x00, 0x30, 0x00,
	0x00, 0xC0, 0x00, 0x03, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x30, 0x00, 0x00,
	0xC0, 0x00, 0x03, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x30, 0x00, 0x00, 0xC0,
	0x00, 0x03, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x30, 0x00, 0x00, 0xC0, 0x00,
	0x03, 0x00, 0x00, 0x0C, 0x00, 0x0F, 0xFF, 0xC0, 0x3F, 0xFF, 0x00, 0xFF,
	0x01, 0xFF, 0xFE, 0x03, 0xFC, 0xC0, 0x00, 0x61, 0x80, 0x00, 0xC3, 0x00,
	0x01, 0x86, 0x00, 0x03, 0x0C, 0x00, 0x06, 0x18, 0x00, 0x0C, 0x30, 0x00,
	0x18, 0x60, 0x00, 0x30, 0xC0, 0x00, 0x61, 0x80, 0x00, 0xC3, 0x00, 0x01,
	0x86, 0x00, 0x03, 0x0C, 0x00, 0x06, 0x18, 0x00, 0x0C, 0x30, 0x00, 0x18,
	0x60, 0x00, 0x30, 0xC0, 0x00, 0x61, 0x80, 0x00, 0xC3, 0x80, 0x03, 0x83,
	0x00, 0x06, 0x07, 0x00, 0x1C, 0x07, 0x00, 0x70, 0x07, 0x83, 0xC0, 0x07,
	0xFF, 0x00, 0x03, 0xF8, 0x00, 0x7F, 0xC0, 0x3F, 0xF7, 0xFC, 0x03, 0xFF,
	0x18, 0x00, 0x01, 0x80, 0xC0, 0x00, 0x30, 0x0C, 0x00, 0x03, 0x00, 0x60,
	0x00, 0x30, 0x06, 0x00, 0x06, 0x00, 0x60, 0x00, 0x60, 0x03, 0x00, 0x0C,
	0x00, 0x30, 0x00, 0xC0, 0x03, 0x80, 0x0C, 0x00, 0x18, 0x01, 0x80, 0x01,
	0x80, 0x18, 0x00, 0x0C, 0x03, 0x00, 0x00, 0xC0, 0x30, 0x00, 0x0E, 0x03,
	0x00, 0x00, 0x60, 0x60, 0x00, 0x06, 0x06, 0x00, 0x00, 0x30, 0xC0, 0x00,
	0x03, 0x0C, 0x00, 0x00, 0x30, 0x80, 0x00, 0x01, 0x98, 0x00, 0x00, 0x19,
	0x80, 0x00, 0x00, 0xF0, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0xE0, 0x00,
	0xFF, 0x80, 0x7F, 0xFF, 0xE0, 0x1F, 0xF3, 0x00, 0x00, 0x30, 0xC0, 0x00,
	0x0C, 0x30, 0x00, 0x03, 0x0C, 0x03, 0x80, 0xC3, 0x01, 0xE0, 0x30, 0x60,
	0x78, 0x0C, 0x18, 0x1F, 0x02, 0x06, 0x04, 0xC0, 0x81, 0x83, 0x30, 0x60,
	0x60, 0xCC, 0x18, 0x18, 0x31, 0x86, 0x06, 0x18, 0x61, 0x81, 0x86, 0x18,
	0x60, 0x71, 0x87, 0x18, 0x0C, 0x40, 0xC6, 0x03, 0x30, 0x31, 0x00, 0xCC,
	0x0C, 0xC0, 0x33, 0x01, 0xB0, 0x0D, 0x80, 0x6C, 0x03, 0x60, 0x1B, 0x00,
	0xD8, 0x06, 0xC0, 0x34, 0x00, 0xF0, 0x07, 0x00, 0x3C, 0x01, 0xC0, 0x0E,
	0x00, 0x7F, 0x00, 0xFF, 0x7F, 0x00, 0xFF, 0x18, 0x00, 0x18, 0x0C, 0x00,
	0x38, 0x0E, 0x00, 0x70, 0x07, 0x00, 0x60, 0x03, 0x00, 0xC0, 0x01, 0x81,
	0x80, 0x01, 0xC3, 0x80, 0x00, 0xE7, 0x00, 0x00, 0x76, 0x00, 0x00, 0x3C,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x7E, 0x00, 0x00, 0x66,
	0x00, 0x00, 0xC3, 0x00, 0x01, 0x81, 0x80, 0x03, 0x81, 0xC0, 0x07, 0x00,
	0xE0, 0x06, 0x00, 0x60, 0x0C, 0x00, 0x30, 0x18, 0x00, 0x18, 0x38, 0x00,
	0x1C, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x00,
	0xFF, 0x18, 0x00, 0x18, 0x0C, 0x00, 0x30, 0x0E, 0x00, 0x70, 0x06, 0x00,
	0x60, 0x03, 0x00, 0xC0, 0x03, 0x81, 0xC0, 0x01, 0x81, 0x80, 0x00, 0xC3,
	0x00, 0x00, 0xE7, 0x00, 0x00, 0x66, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x3C,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x07, 0xFF, 0xE0, 0x07, 0xFF,
	0xE0, 0x7F, 0xFF, 0x9F, 0xFF, 0xE6, 0x00, 0x19, 0x80, 0x0C, 0x60, 0x07,
	0x18, 0x03, 0x86, 0x00, 0xC1, 0x80, 0x70, 0x00, 0x38, 0x00, 0x0C, 0x00,
	0x07, 0x00, 0x03, 0x80, 0x00, 0xC0, 0x00, 0x60, 0x00, 0x38, 0x00, 0x1C,
	0x00, 0x06, 0x00, 0x03, 0x80, 0x31, 0xC0, 0x0C, 0x60, 0x03, 0x30, 0x00,
	0xDC, 0x00, 0x3E, 0x00, 0x0F, 0x00, 0x03, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0,
	0xFF, 0xFF, 0x06, 0x0C, 0x18, 0x30, 0x60, 0xC1, 0x83, 0x06, 0x0C, 0x18,
	0x30, 0x60, 0xC1, 0x83, 0x06, 0x0C, 0x18, 0x30, 0x60, 0xC1, 0x83, 0x06,
	0x0C, 0x18, 0x30, 0x60, 0xFF, 0xFC, 0xC0, 0x00, 0x30, 0x00, 0x06, 0x00,
	0x01, 0x80, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x01, 0x80, 0x00, 0x60, 0x00,
	0x0C, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00,
	0xC0, 0x00, 0x18, 0x00, 0x06, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x06,
	0x00, 0x01, 0x80, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x03, 0x80, 0x00, 0x60,
	0x00, 0x1C, 0x00, 0x03, 0x00, 0x00, 0xE0, 0x00, 0x18, 0x00, 0x07, 0x00,
	0x00, 0xC0, 0x00, 0x30, 0x00, 0x06, 0x00, 0x01, 0x80, 0x00, 0x30, 0x00,
	0x0C, 0xFF, 0xFC, 0x18, 0x30, 0x60, 0xC1, 0x83, 0x06, 0x0C, 0x18, 0x30,
	0x60, 0xC1, 0x83, 0x06, 0x0C, 0x18, 0x30, 0x60, 0xC1, 0x83, 0x06, 0x0C,
	0x18, 0x30, 0x60, 0xC1, 0x83, 0xFF, 0xFC, 0x00, 0x40, 0x00, 0x30, 0x00,
	0x1E, 0x00, 0x0E, 0xC0, 0x07, 0x38, 0x01, 0x87, 0x00, 0xC0, 0xC0, 0x60,
	0x18, 0x38, 0x03, 0x1C, 0x00, 0xE6, 0x00, 0x1F, 0x00, 0x03, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xC0, 0xE0, 0x70, 0x3C, 0x0E, 0x07, 0x03,
	0x01, 0xFC, 0x00, 0x7F, 0xFC, 0x01, 0xC0, 0x3C, 0x00, 0x00, 0x30, 0x00,
	0x00, 0x60, 0x00, 0x01, 0x80, 0x00, 0x06, 0x00, 0x00, 0x18, 0x00, 0x00,
	0x60, 0x0F, 0xF9, 0x81, 0xFF, 0xFE, 0x0F, 0x80, 0x38, 0x70, 0x00, 0x63,
	0x80, 0x01, 0x8C, 0x00, 0x06, 0x30, 0x00, 0x18, 0xC0, 0x00, 0xE3, 0x00,
	0x07, 0x86, 0x00, 0x76, 0x1E, 0x07, 0x9F, 0x3F, 0xF8, 0x7C, 0x3F, 0x80,
	0x00, 0xF8, 0x00, 0x01, 0xF0, 0x00, 0x00, 0x60, 0x00, 0x00, 0xC0, 0x00,
	0x01, 0x80, 0x00, 0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x0C, 0x1F, 0x80,
	0x18, 0xFF, 0xC0, 0x33, 0x81, 0xC0, 0x6E, 0x01, 0xC0, 0xF0, 0x00, 0xC1,
	0xE0, 0x01, 0xC3, 0x80, 0x01, 0x87, 0x00, 0x03, 0x8C, 0x00, 0x03, 0x18,
	0x00, 0x06, 0x30, 0x00, 0x0C, 0x60, 0x00, 0x18, 0xC0, 0x00, 0x31, 0x80,
	0x00, 0x63, 0x80, 0x01, 0x87, 0x00, 0x03, 0x0F, 0x00, 0x0E, 0x1F, 0x00,
	0x38, 0x37, 0x00, 0xE3, 0xE7, 0x03, 0x87, 0xC7, 0xFE, 0x00, 0x03, 0xF0,
	0x00, 0x01, 0xFC, 0x00, 0x3F, 0xF9, 0x83, 0xC0, 0xFC, 0x38, 0x01, 0xE3,
	0x00, 0x07, 0x38, 0x00, 0x19, 0x80, 0x00, 0xDC, 0x00, 0x06, 0xC0, 0x00,
	0x06, 0x00, 0x00, 0x30, 0x00, 0x01, 0x80, 0x00, 0x0C, 0x00, 0x00, 0x60,
	0x00, 0x03, 0x80, 0x00, 0x0C, 0x00, 0x00, 0x70, 0x00, 0x01, 0x80, 0x00,
	0xC7, 0x00, 0x1E, 0x1E, 0x03, 0xC0, 0x7F, 0xFC, 0x00, 0xFF, 0x00, 0x00,
	0x00, 0xF8, 0x00, 0x00, 0xF8, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x01, 0xF8, 0x18, 0x07,
	0xFE, 0x18, 0x0F, 0x07, 0x98, 0x1C, 0x01, 0xD8, 0x38, 0x00, 0xF8, 0x70,
	0x00, 0x78, 0x60, 0x00, 0x38, 0xE0, 0x00, 0x38, 0xC0, 0x00, 0x18, 0xC0,
	0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0,
	0x00, 0x18, 0x60, 0x00, 0x38, 0x60, 0x00, 0x38, 0x70, 0x00, 0x78, 0x38,
	0x00, 0xD8, 0x1C, 0x01, 0xD8, 0x0F, 0x07, 0x9F, 0x07, 0xFE, 0x1F, 0x01,
	0xF8, 0x00, 0x01, 0xFC, 0x00, 0x3F, 0xF8, 0x07, 0x80, 0xF0, 0x70, 0x01,
	0xC3, 0x00, 0x07, 0x30, 0x00, 0x19, 0x80, 0x00, 0x78, 0x00, 0x03, 0xC0,
	0x00, 0x1F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x0C, 0x00, 0x00,
	0x60, 0x00, 0x01, 0x80, 0x00, 0x0C, 0x00, 0x00, 0x30, 0x00, 0x01, 0xC0,
	0x00, 0xC7, 0x00, 0x0E, 0x1E, 0x03, 0xE0, 0x3F, 0xFC, 0x00, 0x7F, 0x00,
	0x00, 0x7F, 0xC0, 0x3F, 0xFC, 0x0E, 0x00, 0x03, 0x80, 0x00, 0x60, 0x00,
	0x0C, 0x00, 0x01, 0x80, 0x00, 0x30, 0x00, 0xFF, 0xFF, 0x9F, 0xFF, 0xF0,
	0x18, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x0C, 0x00, 0x01, 0x80, 0x00,
	0x30, 0x00, 0x06, 0x00, 0x00, 0xC0, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00,
	0x60, 0x00, 0x0C, 0x00, 0x01, 0x80, 0x00, 0x30, 0x00, 0x06, 0x00, 0x00,
	0xC0, 0x03, 0xFF, 0xFC, 0x7F, 0xFF, 0x80, 0x01, 0xF8, 0x00, 0x0F, 0xFC,
	0x7C, 0x38, 0x1C, 0xF8, 0xE0, 0x0D, 0x83, 0x00, 0x0F, 0x0E, 0x00, 0x1E,
	0x18, 0x00, 0x1C, 0x70, 0x00, 0x38, 0xC0, 0x00, 0x31, 0x80, 0x00, 0x63,
	0x00, 0x00, 0xC6, 0x00, 0x01, 0x8C, 0x00, 0x03, 0x18, 0x00, 0x06, 0x18,
	0x00, 0x1C, 0x30, 0x00, 0x38, 0x30, 0x00, 0xF0, 0x70, 0x03, 0x60, 0x78,
	0x1C, 0xC0, 0x3F, 0xF1, 0x80, 0x1F, 0x83, 0x00, 0x00, 0x06, 0x00, 0x00,
	0x0C, 0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0xC0, 0x00, 0x03,
	0x80, 0x00, 0x0E, 0x00, 0x3F, 0xF8, 0x00, 0x7F, 0xC0, 0x00, 0xF8, 0x00,
	0x01, 0xF0, 0x00, 0x00, 0x60, 0x00, 0x00, 0xC0, 0x00, 0x01, 0x80, 0x00,
	0x03, 0x00, 0x00, 0x06, 0x00, 0x00, 0x0C, 0x3F, 0x00, 0x18, 0xFF, 0x80,
	0x37, 0x03, 0x80, 0x7C, 0x03, 0x80, 0xF0, 0x03, 0x81, 0xC0, 0x03, 0x03,
	0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00, 0x30, 0x30,
	0x00, 0x60, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80, 0x03, 0x03, 0x00,
	0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00, 0x30, 0x30, 0x00,
	0x63, 0xFC, 0x07, 0xFF, 0xF8, 0x0F, 0xF0, 0x01, 0xC0, 0x00, 0x70, 0x00,
	0x1C, 0x00, 0x07, 0x00, 0x01, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x0F, 0xF0, 0x03, 0xFC, 0x00, 0x03, 0x00, 0x00, 0xC0,
	0x00, 0x30, 0x00, 0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00,
	0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x03,
	0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x03, 0xFF, 0xFF, 0xFF, 0xFF,
	0xC0, 0x00, 0x70, 0x01, 0xC0, 0x07, 0x00, 0x1C, 0x00, 0x70, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x03, 0xFF, 0xFF, 0xFF, 0xC0, 0x03, 0x00, 0x0C,
	0x00, 0x30, 0x00, 0xC0, 0x03, 0x00, 0x0C, 0x00, 0x30, 0x00, 0xC0, 0x03,
	0x00, 0x0C, 0x00, 0x30, 0x00, 0xC0, 0x03, 0x00, 0x0C, 0x00, 0x30, 0x00,
	0xC0, 0x03, 0x00, 0x0C, 0x00, 0x30, 0x00, 0xC0, 0x03, 0x00, 0x0C, 0x00,
	0x70, 0x03, 0x80, 0x1C, 0xFF, 0xE3, 0xFF, 0x00, 0xF8, 0x00, 0x03, 0xE0,
	0x00, 0x01, 0x80, 0x00, 0x06, 0x00, 0x00, 0x18, 0x00, 0x00, 0x60, 0x00,
	0x01, 0x80, 0x00, 0x06, 0x00, 0x00, 0x18, 0x1F, 0xE0, 0x60, 0x7F, 0x81,
	0x80, 0x60, 0x06, 0x07, 0x00, 0x18, 0x38, 0x00, 0x61, 0xC0, 0x01, 0x8E,
	0x00, 0x06, 0x70, 0x00, 0x1B, 0x80, 0x00, 0x7F, 0x00, 0x01, 0xCE, 0x00,
	0x06, 0x1C, 0x00, 0x18, 0x38, 0x00, 0x60, 0x70, 0x01, 0x80, 0xE0, 0x06,
	0x01, 0xC0, 0x18, 0x03, 0x80, 0x60, 0x07, 0x0F, 0x80, 0x7F, 0xFE, 0x01,
	0xFF, 0x3F, 0xC0, 0x0F, 0xF0, 0x00, 0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0,
	0x00, 0x30, 0x00, 0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00,
	0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x03,
	0x00, 0x00, 0xC0, 0x00, 0x30, 0x00, 0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0,
	0x00, 0x30, 0x00, 0x0C, 0x00, 0x03, 0x00, 0x00, 0xC0, 0x00, 0x30, 0x0F,
	0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0xF0, 0x3C, 0x0F, 0x9F, 0x87, 0xE0, 0xFB,
	0x1C, 0xC7, 0x01, 0xE0, 0xD8, 0x38, 0x1C, 0x07, 0x01, 0x81, 0x80, 0x60,
	0x18, 0x18, 0x06, 0x01, 0x81, 0x80, 0x60, 0x18, 0x18, 0x06, 0x01, 0x81,
	0x80, 0x60, 0x18, 0x18, 0x06, 0x01, 0x81, 0x80, 0x60, 0x18, 0x18, 0x06,
	0x01, 0x81, 0x80, 0x60, 0x18, 0x18, 0x06, 0x01, 0x81, 0x80, 0x60, 0x18,
	0x18, 0x06, 0x01, 0x81, 0x80, 0x60, 0x18, 0x18, 0x06, 0x01, 0x8F, 0xE0,
	0x7C, 0x1F, 0xFE, 0x07, 0xC1, 0xF0, 0x00, 0x1F, 0x00, 0xF8, 0xFF, 0x81,
	0xF3, 0x83, 0x80, 0x6C, 0x03, 0x80, 0xF0, 0x03, 0x81, 0xC0, 0x03, 0x03,
	0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00, 0x30, 0x30,
	0x00, 0x60, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80, 0x03, 0x03, 0x00,
	0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00, 0x30, 0x30, 0x00,
	0x67, 0xFC, 0x03, 0xFF, 0xF8, 0x07, 0xE0, 0x00, 0xFC, 0x00, 0x1F, 0xFE,
	0x00, 0xF0, 0x3C, 0x07, 0x00, 0x38, 0x38, 0x00, 0x71, 0xC0, 0x00, 0xE6,
	0x00, 0x01, 0x98, 0x00, 0x06, 0xC0, 0x00, 0x0F, 0x00, 0x00, 0x3C, 0x00,
	0x00, 0xF0, 0x00, 0x03, 0xC0, 0x00, 0x0F, 0x00, 0x00, 0x36, 0x00, 0x01,
	0x98, 0x00, 0x06, 0x70, 0x00, 0x38, 0xE0, 0x01, 0xC1, 0xC0, 0x0E, 0x03,
	0xC0, 0xF0, 0x07, 0xFF, 0x80, 0x03, 0xF0, 0x00, 0x00, 0x3F, 0x01, 0xF1,
	0xFF, 0x83, 0xE7, 0x03, 0x80, 0xD8, 0x01, 0x81, 0xE0, 0x01, 0x83, 0xC0,
	0x03, 0x87, 0x00, 0x03, 0x0E, 0x00, 0x07, 0x18, 0x00, 0x06, 0x30, 0x00,
	0x0C, 0x60, 0x00, 0x18, 0xC0, 0x00, 0x31, 0x80, 0x00, 0x63, 0x00, 0x00,
	0xC7, 0x00, 0x03, 0x0E, 0x00, 0x06, 0x1E, 0x00, 0x18, 0x36, 0x00, 0x70,
	0x67, 0x03, 0xC0, 0xC7, 0xFE, 0x01, 0x83, 0xF0, 0x03, 0x00, 0x00, 0x06,
	0x00, 0x00, 0x0C, 0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0x60,
	0x00, 0x00, 0xC0, 0x00, 0x0F, 0xFC, 0x00, 0x1F, 0xF8, 0x00, 0x00, 0x01,
	0xF8, 0x00, 0x07, 0xFF, 0x1F, 0x0F, 0x07, 0x9F, 0x1C, 0x01, 0xD8, 0x38,
	0x00, 0x78, 0x70, 0x00, 0x78, 0x60, 0x00, 0x38, 0xE0, 0x00, 0x38, 0xC0,
	0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0, 0x00, 0x18, 0xC0,
	0x00, 0x18, 0xC0, 0x00, 0x18, 0x60, 0x00, 0x38, 0x70, 0x00, 0x78, 0x30,
	0x00, 0x78, 0x1C, 0x01, 0xD8, 0x0F, 0x07, 0x98, 0x07, 0xFF, 0x18, 0x01,
	0xFC, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00,
	0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18, 0x00,
	0x03, 0xFF, 0x00, 0x03, 0xFF, 0x7E, 0x03, 0xC3, 0xF0, 0x7F, 0x81, 0x8F,
	0x0E, 0x0C, 0xE0, 0x00, 0x7E, 0x00, 0x03, 0xC0, 0x00, 0x1C, 0x00, 0x00,
	0xC0, 0x00, 0x06, 0x00, 0x00, 0x30, 0x00, 0x01, 0x80, 0x00, 0x0C, 0x00,
	0x00, 0x60, 0x00, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0xC0, 0x00, 0x06,
	0x00, 0x00, 0x30, 0x00, 0x3F, 0xFF, 0xC1, 0xFF, 0xFE, 0x00, 0x07, 0xF0,
	0x07, 0xFF, 0x63, 0xC0, 0xF9, 0xC0, 0x0E, 0x60, 0x01, 0x98, 0x00, 0x66,
	0x00, 0x19, 0xC0, 0x00, 0x38, 0x00, 0x07, 0xC0, 0x00, 0x7F, 0xC0, 0x00,
	0x7C, 0x00, 0x03, 0x80, 0x00, 0x70, 0x00, 0x0F, 0x00, 0x03, 0xC0, 0x00,
	0xF8, 0x00, 0x7F, 0x00, 0x3B, 0xF0, 0x3C, 0xDF, 0xFE, 0x00, 0xFE, 0x00,
	0x0C, 0x00, 0x00, 0x60, 0x00, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0xC0,
	0x00, 0x06, 0x00, 0x03, 0xFF, 0xFE, 0x1F, 0xFF, 0xF0, 0x0C, 0x00, 0x00,
	0x60, 0x00, 0x03, 0x00, 0x00, 0x18, 0x00, 0x00, 0xC0, 0x00, 0x06, 0x00,
	0x00, 0x30, 0x00, 0x01, 0x80, 0x00, 0x0C, 0x00, 0x00, 0x60, 0x00, 0x03,
	0x00, 0x00, 0x18, 0x00, 0x00, 0xC0, 0x00, 0x06, 0x00, 0x00, 0x30, 0x00,
	0x00, 0xC0, 0x07, 0x07, 0x01, 0xF0, 0x1F, 0xFF, 0x00, 0x3F, 0x80, 0xF8,
	0x03, 0xF1, 0xF0, 0x07, 0xE0, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80,
	0x03, 0x03, 0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x18, 0x18, 0x00,
	0x30, 0x30, 0x00, 0x60, 0x60, 0x00, 0xC0, 0xC0, 0x01, 0x81, 0x80, 0x03,
	0x03, 0x00, 0x06, 0x06, 0x00, 0x0C, 0x0C, 0x00, 0x38, 0x18, 0x00, 0xF0,
	0x18, 0x03, 0x60, 0x38, 0x3C, 0xF8, 0x3F, 0xF1, 0xF0, 0x1F, 0x00, 0x00,
	0x7F, 0xC0, 0xFF, 0xDF, 0xF0, 0x3F, 0xF0, 0xC0, 0x00, 0xC0, 0x30, 0x00,
	0x30, 0x06, 0x00, 0x1C, 0x01, 0x80, 0x06, 0x00, 0x30, 0x01, 0x80, 0x0C,
	0x00, 0xC0, 0x03, 0x80, 0x30, 0x00, 0x60, 0x18, 0x00, 0x18, 0x06, 0x00,
	0x03, 0x03, 0x00, 0x00, 0xC0, 0xC0, 0x00, 0x18, 0x30, 0x00, 0x06, 0x18,
	0x00, 0x00, 0xC6, 0x00, 0x00, 0x33, 0x00, 0x00, 0x0E, 0xC0, 0x00, 0x01,
	0xE0, 0x00, 0x00, 0x78, 0x00, 0x7F, 0x00, 0x3F, 0xDF, 0xC0, 0x0F, 0xF1,
	0x80, 0x00, 0x20, 0x60, 0x00, 0x18, 0x18, 0x00, 0x06, 0x06, 0x03, 0x01,
	0x80, 0x81, 0xE0, 0x60, 0x30, 0x78, 0x10, 0x0C, 0x1E, 0x0C, 0x03, 0x0C,
	0xC3, 0x00, 0xC3, 0x30, 0xC0, 0x10, 0xCC, 0x30, 0x06, 0x61, 0x98, 0x01,
	0x98, 0x66, 0x00, 0x66, 0x19, 0x80, 0x0B, 0x03, 0x60, 0x03, 0xC0, 0xD0,
	0x00, 0xF0, 0x1C, 0x00, 0x38, 0x07, 0x00, 0x0E, 0x01, 0xC0, 0x3F, 0x81,
	0xFE, 0x3F, 0x81, 0xFE, 0x0C, 0x00, 0x38, 0x06, 0x00, 0x70, 0x03, 0x00,
	0xE0, 0x01, 0x81, 0xC0, 0x00, 0xC3, 0x80, 0x00, 0x67, 0x00, 0x00, 0x3C,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x67, 0x00, 0x00, 0xC3,
	0x80, 0x01, 0x81, 0xC0, 0x03, 0x00, 0xE0, 0x06, 0x00, 0x70, 0x0C, 0x00,
	0x38, 0x18, 0x00, 0x1C, 0x7F, 0x81, 0xFF, 0x7F, 0x81, 0xFF, 0x7F, 0x00,
	0xFF, 0x7F, 0x00, 0xFF, 0x18, 0x00, 0x0C, 0x18, 0x00, 0x18, 0x0C, 0x00,
	0x18, 0x0C, 0x00, 0x30, 0x06, 0x00, 0x30, 0x06, 0x00, 0x60, 0x03, 0x00,
	0x60, 0x03, 0x00, 0xC0, 0x01, 0x80, 0xC0, 0x01, 0x81, 0x80, 0x00, 0xC1,
	0x80, 0x00, 0xC3, 0x00, 0x00, 0x63, 0x00, 0x00, 0x66, 0x00, 0x00, 0x36,
	0x00, 0x00, 0x34, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x18, 0x00, 0x00, 0x18,
	0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0x30, 0x00, 0x00, 0x60,
	0x00, 0x00, 0x60, 0x00, 0x00, 0xC0, 0x00, 0x7F, 0xFC, 0x00, 0x7F, 0xFC,
	0x00, 0xFF, 0xFF, 0x7F, 0xFF, 0xB0, 0x01, 0x98, 0x01, 0xCC, 0x01, 0xC0,
	0x00, 0xC0, 0x00, 0xC0, 0x00, 0xC0, 0x00, 0xC0, 0x00, 0xC0, 0x00, 0xE0,
	0x00, 0x60, 0x00, 0x60, 0x00, 0x60, 0x00, 0x60, 0x00, 0x60, 0x03, 0x70,
	0x01, 0xB0, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x00, 0xE0, 0x7C, 0x0C,
	0x03, 0x00, 0x60, 0x0C, 0x01, 0x80, 0x30, 0x06, 0x00, 0xC0, 0x18, 0x03,
	0x00, 0x60, 0x0C, 0x03, 0x00, 0xE0, 0xF0, 0x1E, 0x00, 0x70, 0x06, 0x00,
	0x60, 0x0C, 0x01, 0x80, 0x30, 0x06, 0x00, 0xC0, 0x18, 0x03, 0x00, 0x60,
	0x0C, 0x01, 0x80, 0x18, 0x03, 0xE0, 0x1C, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xF0, 0xE0, 0x1F, 0x00, 0x60, 0x06, 0x00, 0xC0, 0x18,
	0x03, 0x00, 0x60, 0x0C, 0x01, 0x80, 0x30, 0x06, 0x00, 0xC0, 0x18, 0x01,
	0x80, 0x38, 0x01, 0xE0, 0x3C, 0x1C, 0x03, 0x00, 0xC0, 0x18, 0x03, 0x00,
	0x60, 0x0C, 0x01, 0x80, 0x30, 0x06, 0x00, 0xC0, 0x18, 0x03, 0x00, 0xC0,
	0xF8, 0x1C, 0x00, 0x0F, 0x00, 0x03, 0xFC, 0x03, 0x70, 0xE0, 0x76, 0x07,
	0x8E, 0xC0, 0x1F, 0xC0, 0x00, 0xF0}

var Regular24pt7bGlyphs = []tinyfont.Glyph{
	{0, 0, 0, 28, 0, 1},        // 0x20 ' '
	{0, 5, 30, 28, 11, -28},    // 0x21 '!'
	{19, 16, 14, 28, 6, -28},   // 0x22 '"'
	{47, 19, 32, 28, 4, -29},   // 0x23 '#'
	{123, 18, 33, 28, 5, -29},  // 0x24 '$'
	{198, 20, 29, 28, 4, -27},  // 0x25 '%'
	{271, 18, 25, 28, 5, -23},  // 0x26 '&'
	{328, 7, 14, 28, 11, -28},  // 0x27 '''
	{341, 7, 34, 28, 14, -27},  // 0x28 '('
	{371, 7, 34, 28, 8, -27},   // 0x29 ')'
	{401, 18, 16, 28, 5, -27},  // 0x2A '*'
	{437, 20, 22, 28, 4, -23},  // 0x2B '+'
	{492, 9, 14, 28, 6, -6},    // 0x2C ','
	{508, 22, 2, 28, 3, -13},   // 0x2D '-'
	{514, 7, 6, 28, 11, -4},    // 0x2E '.'
	{520, 18, 35, 28, 5, -30},  // 0x2F '/'
	{599, 18, 30, 28, 5, -28},  // 0x30 '0'
	{667, 16, 29, 28, 6, -28},  // 0x31 '1'
	{725, 18, 29, 28, 5, -28},  // 0x32 '2'
	{791, 19, 30, 28, 5, -28},  // 0x33 '3'
	{863, 16, 28, 28, 6, -27},  // 0x34 '4'
	{919, 19, 29, 28, 5, -27},  // 0x35 '5'
	{988, 18, 30, 28, 6, -28},  // 0x36 '6'
	{1056, 18, 28, 28, 5, -27}, // 0x37 '7'
	{1119, 18, 30, 28, 5, -28}, // 0x38 '8'
	{1187, 18, 30, 28, 6, -28}, // 0x39 '9'
	{1255, 7, 21, 28, 11, -19}, // 0x3A ':'
	{1274, 10, 27, 28, 7, -19}, // 0x3B ';'
	{1308, 22, 22, 28, 3, -23}, // 0x3C '<'
	{1369, 24, 9, 28, 2, -17},  // 0x3D '='
	{1396, 21, 22, 28, 4, -23}, // 0x3E '>'
	{1454, 17, 28, 28, 6, -26}, // 0x3F '?'
	{1514, 18, 32, 28, 5, -28}, // 0x40 '@'
	{1586, 28, 26, 28, 0, -25}, // 0x41 'A'
	{1677, 22, 26, 28, 3, -25}, // 0x42 'B'
	{1749, 22, 28, 28, 3, -26}, // 0x43 'C'
	{1826, 22, 26, 28, 3, -25}, // 0x44 'D'
	{1898, 22, 26, 28, 3, -25}, // 0x45 'E'
	{1970, 22, 26, 28, 3, -25}, // 0x46 'F'
	{2042, 23, 28, 28, 3, -26}, // 0x47 'G'
	{2123, 23, 26, 28, 3, -25}, // 0x48 'H'
	{2198, 16, 26, 28, 6, -25}, // 0x49 'I'
	{2250, 23, 27, 28, 4, -25}, // 0x4A 'J'
	{2328, 24, 26, 28, 3, -25}, // 0x4B 'K'
	{2406, 21, 26, 28, 4, -25}, // 0x4C 'L'
	{2475, 26, 26, 28, 1, -25}, // 0x4D 'M'
	{2560, 24, 26, 28, 2, -25}, // 0x4E 'N'
	{2638, 24, 28, 28, 2, -26}, // 0x4F 'O'
	{2722, 21, 26, 28, 3, -25}, // 0x50 'P'
	{2791, 24, 32, 28, 2, -26}, // 0x51 'Q'
	{2887, 24, 26, 28, 3, -25}, // 0x52 'R'
	{2965, 20, 28, 28, 4, -26}, // 0x53 'S'
	{3035, 22, 26, 28, 3, -25}, // 0x54 'T'
	{3107, 23, 27, 28, 3, -25}, // 0x55 'U'
	{3185, 28, 26, 28, 0, -25}, // 0x56 'V'
	{3276, 26, 26, 28, 1, -25}, // 0x57 'W'
	{3361, 24, 26, 28, 2, -25}, // 0x58 'X'
	{3439, 24, 26, 28, 2, -25}, // 0x59 'Y'
	{3517, 18, 26, 28, 5, -25}, // 0x5A 'Z'
	{3576, 7, 34, 28, 13, -27}, // 0x5B '['
	{3606, 18, 35, 28, 5, -30}, // 0x5C '\'
	{3685, 7, 34, 28, 8, -27},  // 0x5D ']'
	{3715, 18, 12, 28, 5, -28}, // 0x5E '^'
	{3742, 28, 2, 28, 0, 5},    // 0x5F '_'
	{3749, 8, 7, 28, 7, -29},   // 0x60 '`'
	{3756, 22, 22, 28, 3, -20}, // 0x61 'a'
	{3817, 23, 29, 28, 2, -27}, // 0x62 'b'
	{3901, 21, 22, 28, 4, -20}, // 0x63 'c'
	{3959, 24, 29, 28, 3, -27}, // 0x64 'd'
	{4046, 21, 22, 28, 3, -20}, // 0x65 'e'
	{4104, 19, 28, 28, 6, -27}, // 0x66 'f'
	{4171, 23, 30, 28, 3, -20}, // 0x67 'g'
	{4258, 23, 28, 28, 3, -27}, // 0x68 'h'
	{4339, 18, 29, 28, 5, -28}, // 0x69 'i'
	{4405, 14, 38, 28, 6, -28}, // 0x6A 'j'
	{4472, 22, 28, 28, 4, -27}, // 0x6B 'k'
	{4549, 18, 28, 28, 5, -27}, // 0x6C 'l'
	{4612, 28, 21, 28, 0, -20}, // 0x6D 'm'
	{4686, 23, 21, 28, 2, -20}, // 0x6E 'n'
	{4747, 22, 22, 28, 3, -20}, // 0x6F 'o'
	{4808, 23, 30, 28, 2, -20}, // 0x70 'p'
	{4895, 24, 30, 28, 3, -20}, // 0x71 'q'
	{4985, 21, 20, 28, 5, -19}, // 0x72 'r'
	{5038, 18, 22, 28, 5, -20}, // 0x73 's'
	{5088, 21, 27, 28, 3, -25}, // 0x74 't'
	{5159, 23, 21, 28, 3, -19}, // 0x75 'u'
	{5220, 26, 20, 28, 1, -19}, // 0x76 'v'
	{5285, 26, 20, 28, 1, -19}, // 0x77 'w'
	{5350, 24, 20, 28, 2, -19}, // 0x78 'x'
	{5410, 24, 29, 28, 2, -19}, // 0x79 'y'
	{5497, 17, 20, 28, 6, -19}, // 0x7A 'z'
	{5540, 11, 34, 28, 8, -27}, // 0x7B '{'
	{5587, 2, 34, 28, 13, -27}, // 0x7C '|'
	{5596, 11, 34, 28, 9, -27}, // 0x7D '}'
	{5643, 20, 6, 28, 4, -15}}  // 0x7E '~'

var Regular24pt7b = tinyfont.Font{
	Regular24pt7bBitmaps,
	Regular24pt7bGlyphs,
	0x20, 0x7E, 47}

// Approx. 6330 bytes
