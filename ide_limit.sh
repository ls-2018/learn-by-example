find / |grep 'idea.properties' |xargs -I F sed -i 's/idea.max.intellisense.filesize=2500/idea.max.intellisense.filesize=5000000/g' F
find / |grep 'idea.properties' |xargs -I F sed -i 's/idea.max.content.load.filesize=20000/idea.max.content.load.filesize=5000000/g' F
find / |grep 'idea.properties'|xargs -I F cat F | grep filesize
