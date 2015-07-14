import QtQuick 2.0

Text {
	id: txt1
	anchors.centerIn: parent

	text: golang.data

	color: "blue"
	font.bold: true
	font.pointSize: 20

	MouseArea{
		anchors.fill: parent
		onClicked:{
			golang.setArg()
		}
	}
}
