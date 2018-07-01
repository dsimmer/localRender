#LocalRender

Basically just a chron that looks for outputs from a command line job and keeps running until they are done

// If we use the Maya installation, render.exe needs to be on the path and use the below command
// Render.exe -r redshift -gpu {0} <scene_file>
    //c:\\Program Files\\Autodesk\\maya2010\\Maya.app\\Contents\\bin\\
    //–s 1 –e 125 for frame start end if necessary
    -proj might be required?
todo:
Instead of dying, recovers gracefully
then copy your folder via FTP or locally (locally for now) or to google drive
see schedules via endpoint