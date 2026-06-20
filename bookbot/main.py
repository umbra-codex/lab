import sys

from stats import (
    calculate_number_of_characters,
    calculate_number_of_words,
    chars_dict_to_sorted_list,
)


def get_book_text(path_to_file):
    with open(path_to_file) as file:
        file_contents = file.read()
    return file_contents


def main():
    if len(sys.argv) == 1:
        print("Usage: python3 main.py <path_to_book>")
        sys.exit(1)

    path = sys.argv[1]
    book = get_book_text(path)
    word_count = calculate_number_of_words(book)
    character_count = calculate_number_of_characters(book)
    sorted_list = chars_dict_to_sorted_list(character_count)

    print("============ BOOKBOT ============")
    print(f"Analyzing book found at {path}...")
    print("----------- Word Count ----------")
    print(f"Found {word_count} total words")
    for item in sorted_list:
        print(f"('{item[0]}', {item[1]})")
    print("============= END ===============")


main()
