#include "stdio.h"
#include "string.h"
#include "stdlib.h"

#define BUFFSIZE 1024
#define TRUE 1
#define FALSE 0

typedef enum {
    RED,
    GREEN,
    BLUE,
    NONE
} Color;

int str_includes(const char *str, const char *find) {
    return strstr(str, find) != NULL;
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
    char *start = strchr(line, ':');
    if (start == NULL) {
        printf("error finding : in the line to start parsing");
        exit(1);
    }

    int max_red = 1, max_green = 1, max_blue = 1;

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

            if (color == RED && count > max_red) {
                max_red = count;
            }

            if (color == GREEN && count > max_green) {
                max_green = count;
            }

            if (color == BLUE && count > max_blue) {
                max_blue = count;
            }

            count_str = strtok_r(NULL, count_delim, &next_count);
        }
        section = strtok_r(NULL, section_delim, &next_section);
    }

    return max_red * max_green * max_blue;
}

int main(int argc, char* argv[]) {
    const char* input_file = "./input.txt";
    FILE* f = fopen(input_file, "r");

    if (f == NULL) {
        printf("error reading from file\n");
        return 1;
    }

    int min_power_set_sum = 0;
    char buffer[BUFFSIZE];
    while (fgets(buffer, sizeof buffer, f) != NULL) {
        printf("parsing line: %s", buffer);
        int n = handle_line(buffer);
        printf("min power set of game is %d\n", n);
        min_power_set_sum += n;
    }

    printf("***\nsum of minimap power sets : %d\n***\n", min_power_set_sum);

    fclose(f);
}