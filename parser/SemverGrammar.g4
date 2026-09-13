lexer grammar SemverGrammar;

/*
 * ANTLR grammar for the Semver 2.0 specification. This is directly translated
 * from the BNF Grammar at:
 * https://semver.org/#backusnaur-form-grammar-for-valid-semver-versions
 *
 * This lexer grammar can be imported into other ANTLR grammars.
 */

/*
 * <valid semver> ::= <version core>
 *                  | <version core> "-" <prerelease>
 *                  | <version core> "+" <build>
 *                  | <version core> "-" <prerelease> "+" <build>
 */
fragment VALID_SEMVER
   : VERSION_CORE ('-' PRERELEASE)? ('+' BUILD)?
   ;

/*
 * <version core> ::= <major> "." <minor> "." <patch>
 */
fragment VERSION_CORE
   : MAJOR '.' MINOR '.' PATCH
   ;

/*
 * <major> ::= <numeric identifier>
 */
fragment MAJOR
   : NUMERIC_IDENTIFIER
   ;

/*
 * <minor> ::= <numeric identifier>
 */
fragment MINOR
   : NUMERIC_IDENTIFIER
   ;

/*
 * <patch> ::= <numeric identifier>
 */
fragment PATCH
   : NUMERIC_IDENTIFIER
   ;

/*
 * <prerelease> ::= <dot separated prerelease identifiers>
 */
fragment PRERELEASE
   : DOT_SEPARATED_PRERELEASE_IDENTIFIERS
   ;

/*
 * <dot separated prerelease identifiers> ::= <prerelease identifier>
 *                                           | <prerelease identifier> "." <dot separated prerelease identifiers>
 */
fragment DOT_SEPARATED_PRERELEASE_IDENTIFIERS
   : PRERELEASE_IDENTIFIER ('.' PRERELEASE_IDENTIFIER)*
   ;

/*
 * <build> ::= <dot separated build identifiers>
 */
fragment BUILD
   : DOT_SEPARATED_BUILD_IDENTIFIERS
   ;

/*
 * <dot separated build identifiers> ::= <build identifier>
 *                                      | <build identifier> "." <dot separated build identifiers>
 */
fragment DOT_SEPARATED_BUILD_IDENTIFIERS
   : BUILD_IDENTIFIER ('.' BUILD_IDENTIFIER)*
   ;

/*
 * <prerelease identifier> ::= <alphanumeric identifier>
 *                           | <numeric identifier>
 */
fragment PRERELEASE_IDENTIFIER
   : ALPHANUMERIC_IDENTIFIER
   | NUMERIC_IDENTIFIER
   ;

/*
 * <build identifier> ::= <alphanumeric identifier>
 *                      | <digits>
 */
fragment BUILD_IDENTIFIER
   : ALPHANUMERIC_IDENTIFIER
   | DIGITS
   ;

/*
 * <alphanumeric identifier> ::= <non digit>
 *                             | <non digit> <identifier characters>
 *                             | <identifier characters> <non digit>
 *                             | <identifier characters> <non digit> <identifier characters>
 */
fragment ALPHANUMERIC_IDENTIFIER
   : NON_DIGIT IDENTIFIER_CHARACTERS?
   | IDENTIFIER_CHARACTERS* NON_DIGIT IDENTIFIER_CHARACTERS*
   ;

/*
 * <numeric identifier> ::= "0"
 *                        | <positive digit>
 *                        | <positive digit> <digits>
 */
fragment NUMERIC_IDENTIFIER
   : '0'
   | POSITIVE_DIGIT DIGITS?
   ;

/*
 * <identifier characters> ::= <identifier character>
 *                           | <identifier character> <identifier characters>
 */
fragment IDENTIFIER_CHARACTERS
   : IDENTIFIER_CHARACTER+
   ;

/*
 * <identifier character> ::= <digit>
 *                          | <non digit>
 */
fragment IDENTIFIER_CHARACTER
   : DIGIT
   | NON_DIGIT
   ;

/*
 * <non digit> ::= <letter>
 *               | "-"
 */
fragment NON_DIGIT
   : LETTER
   | '-'
   ;

/*
 * <digits> ::= <digit>
 *            | <digit> <digits>
 */
fragment DIGITS
   : DIGIT+
   ;

/*
 * <digit> ::= "0"
 *           | <positive digit>
 */
fragment DIGIT
   : '0'
   | POSITIVE_DIGIT
   ;

/*
 * <positive digit> ::= "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
 */
fragment POSITIVE_DIGIT
   : [1-9]
   ;

/*
 * <letter> ::= "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J"
 *            | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T"
 *            | "U" | "V" | "W" | "X" | "Y" | "Z" | "a" | "b" | "c" | "d"
 *            | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n"
 *            | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x"
 *            | "y" | "z"
 */
fragment LETTER
   : [A-Za-z]
   ;
