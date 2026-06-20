def calculate_number_of_words(book):
    separated_text = book.split()
    count = 0
    for text in separated_text:
        count += 1
    return count

def calculate_number_of_characters(book):
    separated_text = book.lower()
    char_count = {}
    for char in separated_text:
        if char not in char_count:
            char_count[char] = 1
        elif char in char_count:
            char_count[char] += 1
    return char_count

def chars_dict_to_sorted_list(chars_dict):
  sorted_list = []
  for char in chars_dict:
    if char.isalpha():
      sorted_list.append((char, chars_dict[char]))
  sorted_list = sorted(sorted_list, key=sort_on, reverse=True)
  return sorted_list

def sort_on(items: tuple[str, int]) -> int:
    return items[1]