package main

import "testing"

func TestGetARecords_Empty_Dump_Should_Return_0_Record(t *testing.T) {
	expected := 0
	actual := 0
	dumpFile := "./test_files/cache_dump.0.db"
	aRecords := GetARecords(&dumpFile)
	actual = len(aRecords)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestGetARecords_One_Record_Dump_Should_Return_1_Record(t *testing.T) {
	expected := 1
	actual := 0
	dumpFile := "./test_files/cache_dump.1.db"
	aRecords := GetARecords(&dumpFile)
	actual = len(aRecords)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestGetARecords_Typical_Dump_Should_Return_51_Records(t *testing.T) {
	expected := 51
	actual := 0
	dumpFile := "./test_files/cache_dump.2.db"
	aRecords := GetARecords(&dumpFile)
	actual = len(aRecords)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestGetARecords_Commented_Dump_Should_Return_0_Records(t *testing.T) {
	expected := 0
	actual := 0
	dumpFile := "./test_files/cache_dump.3.db"
	aRecords := GetARecords(&dumpFile)
	actual = len(aRecords)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
