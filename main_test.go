package main

import (
	"testing"
)

func TestMoveUp(t *testing.T) {
	var actual Field
	var expected Field
	actual.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 2, 0, 0, 4, -1},
		{-1, 2, 0, 0, 4, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	expected.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 4, 0, 0, 8, -1},
		{-1, 4, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	actual.up()

	if (actual.matrix != expected.matrix) {
		show(expected)
		show(actual)
		t.Error()
	}
}

func TestMoveDown(t *testing.T) {
	var actual Field
	var expected Field
	actual.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	expected.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 0, 0, 0, 0, -1},
		{-1, 0, 0, 0, 0, -1},
		{-1, 4, 0, 0, 0, -1},
		{-1, 4, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	actual.down()

	if (actual.matrix != expected.matrix) {
		show(expected)
		show(actual)
		t.Error()
	}
}

func TestMoveRight(t *testing.T) {
	var actual Field
	var expected Field
	actual.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, 2, 0, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	expected.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 0, 0, 0, 2, -1},
		{-1, 0, 0, 0, 2, -1},
		{-1, 0, 0, 0, 2, -1},
		{-1, 0, 0, 0, 2, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	actual.right()

	if (actual.matrix != expected.matrix) {
		show(expected)
		show(actual)
		t.Error()
	}
}

func TestMoveLeft(t *testing.T) {
	var actual Field
	var expected Field
	actual.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 2, 0, 2, 0, -1},
		{-1, 2, 0, 2, 0, -1},
		{-1, 2, 0, 4, 0, -1},
		{-1, 2, 0, 8, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	expected.matrix = [FIELD_HEIGHT][FIELD_WIDTH]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, 4, 0, 0, 0, -1},
		{-1, 4, 0, 0, 0, -1},
		{-1, 2, 4, 0, 0, -1},
		{-1, 2, 8, 0, 0, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	actual.left()

	if (actual.matrix != expected.matrix) {
		show(expected)
		show(actual)
		t.Error()
	}
}