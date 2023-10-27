# Use upload.sh to upload the content of ./pages
# This will instantly update the content on thalkz.com
# without needing to make a deployment of the server

scp -r ./pages/* ubuntu@thalkz.com:/home/ubuntu/thalkz.com/pages
