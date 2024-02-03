package testutils

import "os"

func FindTestDbAndRemove(dbName string) error {
	_, err := os.Stat(dbName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if err = os.Remove(dbName); err != nil {
		return err
	}

	return nil
}
