#cd /home/student14
#find -name "*.sh"|cut -d'/' -f2|cut -d"." -f1
find -name "*.sh" -exec basename {} .sh \; | sort -r