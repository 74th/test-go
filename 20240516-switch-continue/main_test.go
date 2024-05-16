package main

import "testing"

func TestSwitchBreak(t *testing.T) {

	switchFlag := 10
	checkPoint1 := false

	for {
		switch switchFlag {
		case 10:
			// for内switch内break
			break
		default:
			t.Error("switch内でbreakするとdefaultは実行されない")
		}
		checkPoint1 = true

		break
	}

	if !checkPoint1 {
		t.Error("switch内でbreakしてもforはbreakしない")
	}
}

func TestSwitchContinue(t *testing.T) {

	switchFlag := 10
	checkPoint1 := false

	for {
		if checkPoint1 {
			break
		}

		checkPoint1 = true

		switch switchFlag {
		case 10:
			// for内switch内continue
			continue
		default:
			t.Error("switch内でcontinueすると、forのcontinueとして働き、defaultは実行されない")
		}

		t.Error("switch内でcontinueすると、forのcontinueとして働き、defaultは実行されない")

		break
	}
}
