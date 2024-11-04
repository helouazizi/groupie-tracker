#include <stdio.h>
#include <string.h>

int main() {
    // Open a file using filestream
    FILE *file = fopen("example.txt", "w");
    if (file == NULL) {
        perror("Error opening file");
        return 1;
    }

    // Write to the file
    const char *text = "Hello, World!\n";
    size_t bytes_written = fwrite(text, sizeof(char), strlen(text), file);
    if (bytes_written < strlen(text)) {
        perror("Error writing to file");
        fclose(file);
        return 1;
    }

    // Close the filestream
    if (fclose(file) != 0) {
        perror("Error closing file");
        return 1;
    }

    return 0;
}