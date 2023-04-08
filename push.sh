git add .
echo "Give the name of the commit"
read commit
git commit -m "$commit"
git push