package main

/*
t=Integer(gets.chomp)
t.times do
	a=gets.chomp.split(" ").map { |x| x.to_i }
	frame=1
	i=0
	score=0
	while frame<11
		if a[i]==10
			score+=(10+a[i+1]+a[i+2])
			i+=1
		elsif a[i]+a[i+1]==10
			score+=(10+a[i+2])
			i+=2
		else
			score+=(a[i]+a[i+1])
			i+=2
		end
		frame+=1
	end
	puts score
end
*/
