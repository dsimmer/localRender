#LocalRender

Basically just a chron that looks for outputs from a command line job and keeps running until they are done

// If we use the Maya installation, render.exe needs to be on the path and use the below command
// Render.exe -r redshift -gpu {0} <scene_file>

todo:
render out of Google drive
static ip for render trigger https://openwrt.org/ raspberry pi wakes main comp? 
onwakelan?
then copy your folder via FTP or locally (locally for now) or to google drive
hit an online endpoint to schedule job for file name
see schedules via endpoint