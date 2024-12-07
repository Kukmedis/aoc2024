package main

type testFixture struct {
	numbers []int
	result  int
}

type operatorsFunc func(testFixture) bool

func isPassingOperatorTest(test testFixture, operators operatorsFunc) bool {
	numbers := test.numbers
	result := test.result
	if len(numbers) == 1 && (numbers[0] == result) {
		return true
	} else if numbers[0] > result || len(numbers) < 2 {
		return false
	} else {
		return operators(test)
	}
}

func calculateCalibrationResult(tests []testFixture) int {
	sum := 0
	for _, t := range tests {
		if isPassingOperatorTest(t, operatorForSimpleCalibration) {
			sum = sum + t.result
		}
	}
	return sum
}

func calculateExtendedCalibrationResult(tests []testFixture) int {
	sum := 0
	for _, t := range tests {
		if isPassingOperatorTest(t, operatorForExtendedCalibration) {
			sum = sum + t.result
		}
	}
	return sum
}

func concatenate(a int, b int) int {
	divResult := 10 * b
	result := a
	for divResult > 9 {
		result = result * 10
		divResult = divResult / 10
	}
	result = result + b
	return result
}

func operatorForSimpleCalibration(test testFixture) bool {
	numbers := test.numbers
	result := test.result
	addResult := append([]int{numbers[0] + numbers[1]}, numbers[2:]...)
	multiplyResult := append([]int{numbers[0] * numbers[1]}, numbers[2:]...)
	return isPassingOperatorTest(testFixture{addResult, result}, operatorForSimpleCalibration) ||
		isPassingOperatorTest(testFixture{multiplyResult, result}, operatorForSimpleCalibration)
}

func operatorForExtendedCalibration(test testFixture) bool {
	numbers := test.numbers
	result := test.result
	addResult := append([]int{numbers[0] + numbers[1]}, numbers[2:]...)
	multiplyResult := append([]int{numbers[0] * numbers[1]}, numbers[2:]...)
	concatenateResult := append([]int{concatenate(numbers[0], numbers[1])}, numbers[2:]...)
	return isPassingOperatorTest(testFixture{addResult, result}, operatorForExtendedCalibration) ||
		isPassingOperatorTest(testFixture{multiplyResult, result}, operatorForExtendedCalibration) ||
		isPassingOperatorTest(testFixture{concatenateResult, result}, operatorForExtendedCalibration)
}
