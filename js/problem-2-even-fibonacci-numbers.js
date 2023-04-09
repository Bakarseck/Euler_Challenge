function fiboEvenSum(n) {
  let prev=0, next=0, current=1, sum=0
  while (current <= n) {
    if (current%2==0) {
      sum+=current
    }
    next = prev+current
    prev=current
    current=next
  }
  console.log(sum)
  return sum;
}

fiboEvenSum(10)
