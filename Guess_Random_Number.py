# Guessing Game
# File name: Guess_Random_Number.py
# Author: Natchathiram Premkumar
# Date: 16/07/2021, Initial version

from random import randrange, uniform

# randrange gives you an integral value
RandomNo = randrange(1, 100)

# Initializing the number of guesses.
count = 0
MaxCount = 10
Left = MaxCount
while count < MaxCount:
	count += 1
	Left  -= 1

	# taking guessing number as input
	UserInput = int(input("Enter guess what the target number is:- "))

	# Condition testing
	if RandomNo == UserInput:
		print("Good Job! You guessed it!")
		
		# Once guessed, loop will break
		break
		
	elif RandomNo > UserInput:
		print("Oops, Your guess was LOW.")
		print("You have only left",Left,"Chances")
		
	elif RandomNo < UserInput:
		print("Oops, Your guess was HIGH.")
		print("You have only left",Left,"Chances")

if count >= MaxCount:
	print("\nSorry, You didn't guess my number. It was %d" % RandomNo)