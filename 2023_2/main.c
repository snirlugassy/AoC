#include "stdio.h"
#include "string.h"
#include "stdlib.h"

#define BUFFSIZE 1024
#define TRUE 1
#define FALSE 0

// only 12 red cubes, 13 green cubes, and 14 blue cubes
#define MAX_RED 12
#define MAX_GREEN 13
#define MAX_BLUE 14

typedef enum {
    RED,
    GREEN,
    BLUE,
    NONE
} Color;

int is_digit(char ch) {
    return '0' <= ch && ch <= '9';
}

int str_includes(const char *str, const char *find) {
    return strstr(str, find) != NULL;
}

int parse_game_number(char *line) {
    // find position of the first digit
    char *start = line;
    while (!is_digit(*start)) {
        start++;
    }

    // convert to integer, strtol parses the first integer only
    char * end;
    return strtol(start, &end, 10);
}

Color which_color(const char *str) {
    if (str_includes(str, "red")) {
        return RED;
    }

    if (str_includes(str, "green")) {
        return GREEN;
    }

    if (str_includes(str, "blue")) {
        return BLUE;
    }

    return NONE;
}

int handle_line(char *line) {
    int game_number = parse_game_number(line);

    char *start = strchr(line, ':');
    if (start == NULL) {
        printf("error finding : in the line to start parsing");
        exit(1);
    }

    start++;

    // split by ;
    const char *section_delim = ";";
    char *next_section;
    char * section = strtok_r(start, section_delim, &next_section);
    while(section) {
        // split by ,
        const char *count_delim = ",";
        char *next_count;
        const char * count_str = strtok_r(section, count_delim, &next_count);
        while(count_str) {
            char *_end;
            int count = strtol(count_str, &_end, 10);
            if (_end == NULL) {
                printf("failed to parse integer in %s\n", count_str);
                exit(1);
            }

            Color color = which_color(count_str);
            if (color == NONE) {
                printf("failed to parse integer in %s\n", count_str);
                exit(1);
            }

            if (color == RED && count > MAX_RED) {
                return 0;
            }

            if (color == GREEN && count > MAX_GREEN) {
                return 0;
            }

            if (color == BLUE && count > MAX_BLUE) {
                return 0;
            }

            count_str = strtok_r(NULL, count_delim, &next_count);
        }
        section = strtok_r(NULL, section_delim, &next_section);
    }

    return game_number;
}

int main(int argc, char* argv[]) {
    const char* input_file = "./input.txt";
    FILE* f = fopen(input_file, "r");

    if (f == NULL) {
        printf("error reading from file\n");
        return 1;
    }

    int valid_game_number_sum = 0;
    char buffer[BUFFSIZE];
    while (fgets(buffer, sizeof buffer, f) != NULL) {
        printf("parsing line: %s", buffer);
        valid_game_number_sum += handle_line(buffer);
    }

    printf("***\nvalid game numbers sum : %d\n***\n", valid_game_number_sum);

    fclose(f);
}