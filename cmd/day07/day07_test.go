package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPassingOperatorTest(t *testing.T) {
	assert.True(t, isPassingOperatorTest(testFixture{[]int{10, 19}, 190}, operatorForSimpleCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{81, 40, 27}, 3267}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{17, 5}, 83}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{15, 6}, 156}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{6, 8, 6, 15}, 7290}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{16, 10, 13}, 161011}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{17, 8, 14}, 192}, operatorForSimpleCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{9, 7, 18, 13}, 21037}, operatorForSimpleCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{11, 6, 16, 20}, 292}, operatorForSimpleCalibration))
}

func TestIsPassingExtendedOperatorTest(t *testing.T) {
	assert.True(t, isPassingOperatorTest(testFixture{[]int{10, 19}, 190}, operatorForExtendedCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{81, 40, 27}, 3267}, operatorForExtendedCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{17, 5}, 83}, operatorForExtendedCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{15, 6}, 156}, operatorForExtendedCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{6, 8, 6, 15}, 7290}, operatorForExtendedCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{16, 10, 13}, 161011}, operatorForExtendedCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{17, 8, 14}, 192}, operatorForExtendedCalibration))
	assert.False(t, isPassingOperatorTest(testFixture{[]int{9, 7, 18, 13}, 21037}, operatorForExtendedCalibration))
	assert.True(t, isPassingOperatorTest(testFixture{[]int{11, 6, 16, 20}, 292}, operatorForExtendedCalibration))
}

func TestCalculateCalibrationResult(t *testing.T) {
	assert.Equal(t, 3749, calculateCalibrationResult(
		[]testFixture{
			{[]int{10, 19}, 190},
			{[]int{81, 40, 27}, 3267},
			{[]int{17, 5}, 83},
			{[]int{15, 6}, 156},
			{[]int{6, 8, 6, 15}, 7290},
			{[]int{16, 10, 13}, 161011},
			{[]int{17, 8, 14}, 192},
			{[]int{9, 7, 18, 13}, 21037},
			{[]int{11, 6, 16, 20}, 292}}))
}

func TestCalculateExtendedCalibrationResult(t *testing.T) {
	assert.Equal(t, 11387, calculateExtendedCalibrationResult(
		[]testFixture{
			{[]int{10, 19}, 190},
			{[]int{81, 40, 27}, 3267},
			{[]int{17, 5}, 83},
			{[]int{15, 6}, 156},
			{[]int{6, 8, 6, 15}, 7290},
			{[]int{16, 10, 13}, 161011},
			{[]int{17, 8, 14}, 192},
			{[]int{9, 7, 18, 13}, 21037},
			{[]int{11, 6, 16, 20}, 292}}))
}

func TestConcatenate(t *testing.T) {
	assert.Equal(t, 12345, concatenate(12, 345))
}
