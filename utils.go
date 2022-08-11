package restic

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// print is used to output stdout and stderr in real time.
func print(stdout, stderr io.ReadCloser, done chan struct{}) {
	stopCh := make(chan struct{}, 2)
	defer stdout.Close()
	defer stderr.Close()
	//errCh := make(chan error, 2)

	// A goroutine that outputs stdout in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stdout already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stdout error:", err)
			return
		}
	}()

	// A goroutine that outputs stderr in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stderr already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stderr error:", err)
			return
		}
	}()

	<-stopCh
	<-stopCh
	done <- struct{}{}
}

// concat restic command all flags
// "omitempty" tag is valid for type "int", "int32", "int64" and "map[string]string"
func concat(f interface{}) string {
	var s string
	t := reflect.TypeOf(f).Elem()
	v := reflect.ValueOf(f).Elem()

	for i := 0; i < v.NumField(); i++ {
		// PkgPath is the package path that qualified a lows case (unexported)
		// field name. It is empty for upper case (exported) field names.
		// Skip unexported fields
		if len(t.Field(i).PkgPath) != 0 {
			continue
		}
		knd := v.Field(i).Kind()
		typ := v.Field(i).Type()
		val := v.Field(i).Interface()
		nam := t.Field(i).Name
		//nam := v.Type().Field(i).Name
		tag := t.Field(i).Tag.Get("json")

		if !v.IsValid() {
			continue
		}
		if v.IsZero() {
			continue
		}
		// If json tag have multiple values, the frist value separated by "," is
		// restic flag name.
		flagName := strings.Split(tag, ",")[0]
		omitempty := strings.Contains(tag, "omitempty")

		_ = knd
		_ = typ
		_ = val
		_ = nam
		_ = tag

		switch typ.String() {
		case "string":
			flagValue, ok := val.(string)
			if ok {
				if len(flagValue) == 0 {
					continue
				}
				s = s + " " + flagName + "=" + flagValue
				s = strings.TrimSpace(s)
			}
		case "[]string":
			flagValue, ok := val.([]string)
			if ok {
				if flagValue == nil {
					continue
				}
				s = s + " " + flagName + "=" + strings.Join(flagValue, ",")
				s = strings.TrimSpace(s)
			}
		case "int":
			flagValue, ok := val.(int)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.Itoa(flagValue)
				s = strings.TrimSpace(s)
			}
		case "int32":
			flagValue, ok := val.(int32)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.FormatInt(int64(flagValue), 10)
				s = strings.TrimSpace(s)
			}
		case "int64":
			flagValue, ok := val.(int64)
			if ok {
				if flagValue == 0 && omitempty {
					continue
				}
				s = s + " " + flagName + "=" + strconv.FormatInt(flagValue, 10)
				s = strings.TrimSpace(s)
			}
		case "bool":
			flagValue, ok := val.(bool)
			if ok {
				if flagValue == false {
					continue
				}
				s = s + " " + flagName
				s = strings.TrimSpace(s)
			}
		case "map[string]string":
			flagValue, ok := val.(map[string]string)
			if ok {
				if len(flagValue) == 0 && omitempty {
					continue
				}
				var ts string
				for key, val := range flagValue {
					ts = ts + key + "=" + val + ","
				}
				ts = strings.TrimSuffix(ts, ",")
				s = s + " " + flagName + "=" + ts
				s = strings.TrimSpace(s)
			}
		}
	}

	return s
}

// concatFlags will concat restic and it's sub-commands all flags
func concatAll(cl ...Command) string {
	var s string
	for _, c := range cl {
		s = s + c.Flags()
	}
	return s
}
