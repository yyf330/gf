package gfpool_test

import (
	"github.com/yyf330/gf/os/gfile"
	"github.com/yyf330/gf/os/gfpool"
	"github.com/yyf330/gf/os/gtime"
	"github.com/yyf330/gf/test/gtest"
	"github.com/yyf330/gf/text/gstr"
	"os"
	"testing"
)

func Test_ConcurrentOS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.TempDir(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 100; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 100; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2200)
	})

	gtest.C(t, func(t *gtest.T) {
		path := gfile.TempDir(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
	gtest.C(t, func(t *gtest.T) {
		path := gfile.TempDir(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		s1 := ""
		for i := 0; i < 1000; i++ {
			s1 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s1))
		t.Assert(err, nil)

		s2 := ""
		for i := 0; i < 1000; i++ {
			s2 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s2))
		t.Assert(err, nil)

		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
	// DATA RACE
	//gtest.C(t, func(t *gtest.T) {
	//	path := gfile.TempDir(gtime.TimestampNanoStr())
	//	defer gfile.Remove(path)
	//	f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f1.Close()
	//
	//	f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	//})
}

func Test_ConcurrentGFPool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path := gfile.TempDir(gtime.TimestampNanoStr())
		defer gfile.Remove(path)
		f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	})
	// DATA RACE
	//gtest.C(t, func(t *gtest.T) {
	//	path := gfile.TempDir(gtime.TimestampNanoStr())
	//	defer gfile.Remove(path)
	//	f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f1.Close()
	//
	//	f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(gstr.Count(gfile.GetContents(path), "@1234567890#"), 2000)
	//})
}
