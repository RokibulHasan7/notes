hey: one two
	# Outputs "hey", since this is the target name
	echo $@

	# outputs all prerequisites newer than the target
	echo $?

	# outputs all prerequisites
	echo $^

	touch hey

one:
	touch one

two:
	touch two

clean:
	rm -f hey one two