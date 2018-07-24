Go exercises

## First steps

In your web browser, open the Go playground at [https://play.golang.org]. Alternatively
you may want to use an installed Go compiler - installation instructions are
at [https://golang.org/dl]

- Change the text printed to the screen by modifying what appears between quotes.
Have the computer greet you by name.

- Display two lines of text by writing a second line of code within the body {} of the
main function.

Shortcut: you can use Shift+Enter to execute the code without clicking the Run
button.

## Exercise 1: Ticket generator

Your challenge is to write
a ticket generator in the Go Playground that makes use of variables,
constants, switch, if, and for. It should also draw on the fmt and
math/rand packages to display and align text and to generate random
numbers.

When planning a trip to Mars, it would be handy to have ticket
pricing from multiple spacelines in one place. Websites exist that
aggregate ticket prices for airlines, but so far nothing exists for
spacelines. That’s not a problem for you, though. You can use Go to
teach your computer to solve problems like this.

	Spaceline        Days Trip type  Price
	======================================
	Virgin Galactic    23 Round-trip $  96
	Virgin Galactic    39 One-way    $  37
	SpaceX             31 One-way    $  41
	Space Adventures   22 Round-trip $ 100
	Space Adventures   22 One-way    $  50
	Virgin Galactic    30 Round-trip $  84
	Virgin Galactic    24 Round-trip $  94
	Space Adventures   27 One-way    $  44
	Space Adventures   28 Round-trip $  86
	SpaceX             41 Round-trip $  72

The table should have four columns:

- The spaceline company providing the service
- The duration in days for the trip to Mars (one-way)
- Whether the price covers a return trip
- The price in millions of dollars 😦

For each ticket, randomly select one of the following spacelines: Space
Adventures, SpaceX, or Virgin Galactic.

Use October 13, 2020 as the departure date for all tickets. Mars will
be 62,100,000 km away from the Earth at the time.

Randomly choose the speed the ship will travel, from 16 to 30 km/s. This
will determine the duration for the trip to Mars and also the ticket
price. Make faster ships more expensive, ranging in price from $36 to
$50 million. Double the price for round trips.


## Exercise 2: The Vigenère cipher

The Vigenère cipher is a 16th century variant of the Caesar cipher. For this exercise, you will write a program to decipher text using a keyword.

With the Caesar cipher, a plain text message is ciphered by shifting each letter ahead by three. The direction is reversed to decipher the resulting message.

Assign each English letter a numeric value, where A = 0, B = 1, all the way to Z = 25. With this in mind, a shift by 3 can be represented by the letter D (D = 3).

To decipher the text in below, start with the letter L and shift it by D. Because L = 11 and D = 3, the result of 11-3 is 8, or the letter I. Should you need to decipher the letter A, it should wrap around to become X, as you saw in lesson 9.

	LFDPHLVDZLFRQTXHUHG
	DDDDDDDDDDDDDDDDDDD

The Caesar cipher and ROT13 are susceptible to what's called _frequency analysis_. Letters that occur frequently in the English language, such as E, will occur frequently in the ciphered text as well. By looking for patterns in the ciphered text, the code can cracked.

To thwart would-be code crackers, the Vigenère cipher shifts each letter based on a repeating keyword, rather than a constant like 3 or 13. The keyword repeats until the end of the message, as shown for the keyword GOLANG in below.

Now that you know what the Vigenère cipher is, you may notice that Vigenère with the keyword D is equivalent to the Caesar cipher. Likewise, ROT13 has a keyword of N (N = 13). Longer keywords are needed to be of any benefit.

	CSOITEUIWUIZNSROCNKFD
	GOLANGGOLANGGOLANGGOL

Write a program to decipher the ciphered text shown in below. To keep it simple, all characters are uppercase English letters for both the text and keyword.

cipherText := "CSOITEUIWUIZNSROCNKFD"
keyword := "GOLANG"

Notes:

- The `strings.Repeat` function may come in handy. Give it a try, but also complete this exercise without importing any packages other than `fmt` to print the deciphered message.
-  Try this exercise using `range` in a loop and again without it. Remember that the `range` keyword splits a string into runes, whereas an index like `keyword[0]` results in a byte.

- You can only perform operations on values of the same type, but you can convert one type to the other (`string`, `byte`, `rune`).

- To wrap around at the edges of the alphabet, the Caesar cipher exercise made use of a comparison. Solve this exercise without any `if` statements by using modulus (`%`).
If you recall, modulus gives the remainder of dividing two numbers. For example, `27 % 26` is `1`, keeping numbers within the 0-25 range. Be careful with negative numbers, though, as `-3 % 26` is still `-3`.
