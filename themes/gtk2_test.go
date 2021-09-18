/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package themes

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGtk2Infos(t *testing.T) {
	Convey("Test gtk2 infos", t, func(c C) {
		infos := gtk2FileReader("testdata/gtkrc-2.0")
		c.So(len(infos), ShouldEqual, 16)

		info := infos.Get("gtk-theme-name")
		c.So(info.value, ShouldEqual, "\"Paper\"")

		info.value = "\"Deepin\""
		c.So(info.value, ShouldEqual, "\"Deepin\"")

		infos = infos.Add("gtk2-test", "test")
		c.So(len(infos), ShouldEqual, 17)
	})

	Convey("Test nil infos", t, func(c C) {
		var infos = gtk2FileReader("testdata/xxx")
		infos = infos.Add("gtk2-test", "test")
		c.So(len(infos), ShouldEqual, 1)
		info := infos.Get("gtk2-test")
		c.So(info.value, ShouldEqual, "test")

		err := gtk2FileWriter(infos, "testdata/tmp-gtk2rc")
		defer os.Remove("testdata/tmp-gtk2rc")
		c.So(err, ShouldBeNil)
	})
}
