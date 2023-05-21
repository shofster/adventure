@rem  --pull
fyne-cross windows -arch=amd64  -icon=building.png -app-id=com.scsi.adventure
fyne-cross linux -arch=amd64 -icon=building.png -app-id=com.scsi.adventure
fyne-cross linux -arch=arm64 -icon=building.png -app-id=com.scsi.adventure
fyne-cross linux -arch=arm -icon=building.png -app-id=com.scsi.adventure


fyne build
strip adventure.exe