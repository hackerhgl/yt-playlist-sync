# Youtube playlist sync

[yt-dlp](https://github.com/yt-dlp/yt-dlp) cli library is utilized for this project.

## Credentials
- `PLoSjAzdJQCyfgkgOEzP6ZPJlKmb1w2FsD` Testing

## Test scripts

yt-dlp -f bestaudio --embed-thumbnail --add-metadata --audio-quality 0 0EVVKs6DQLo

yt-dlp -f bestaudio --audio-format mp3 --audio-quality 0 0EVVKs6DQLo

yt-dlp -f bestaudio --embed-thumbnail --add-metadata --extract-audio --audio-format mp3 --audio-quality 0 0EVVKs6DQLo

yt-dlp -f bestaudio --embed-thumbnail --add-metadata --extract-audio --audio-format mp3 --audio-quality 0 -o '%(title)s.%(ext)s' 0EVVKs6DQLo


yt-dlp -f bestaudio --extract-audio --audio-format mp3 --audio-quality 0 -o '%(title)s.%(ext)s' 0EVVKs6DQLo
