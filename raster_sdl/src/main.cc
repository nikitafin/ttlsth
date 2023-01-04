//
// Created by ENBSF on 8/16/2022.
//

#include <SDL.h>

constexpr int gWindowWidth = 640;
constexpr int gWindowHeight = 480;

//    SDL_Color oldColor = {};
//    SDL_GetRenderDrawColor(&sdlRenderer, &oldColor.r, &oldColor.g, &oldColor.b, &oldColor.a);

void line(int x0, int y0, int x1, int y1, SDL_Renderer &sdlRenderer, SDL_Color const &color) {
    SDL_SetRenderDrawColor(&sdlRenderer, color.r, color.g, color.b, color.a);
#if 0
    for (float t = 0; t < 1; t += 0.1f) {
        int x = x0 + (x1 - x0) * t;
        int y = y0 + (y1 - y0) * t;
        SDL_RenderDrawPoint(&sdlRenderer, x, y);
    }
#else
    for (int x = x0; x <= x1; ++x) {
        float t = (x - x0) / static_cast<float>(x1 - x0);
        int y = y0 * (1.0f - t) + y1 * t;
        SDL_RenderDrawPoint(&sdlRenderer, x, y);
    }
#endif
}

constexpr SDL_Color red = {255, 0, 0, 255};
constexpr SDL_Color white = {255, 255, 255, 255};

int main(int argc, char *argv[]) {
    SDL_Event event;
    SDL_Renderer *renderer;
    SDL_Window *window;

    SDL_Init(SDL_INIT_VIDEO);
    SDL_CreateWindowAndRenderer(gWindowWidth, gWindowHeight, 0, &window, &renderer);
    SDL_SetRenderDrawColor(renderer, 0, 0, 0, 0);
    SDL_RenderClear(renderer);

#if 0
    line(1, 1, 199, 100, *renderer, SDL_Color{255, 255, 255, 255});
    line(13, 1, 199, 100, *renderer, SDL_Color{255, 255, 255, 255});
#endif
    line(13, 20, 80, 40, *renderer, white);
    line(20, 13, 40, 80, *renderer, red);
    line(80, 40, 13, 20, *renderer, red);

    SDL_RenderPresent(renderer);
    while (true) {
        if (SDL_PollEvent(&event) && event.type == SDL_QUIT)
            break;
    }
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);
    SDL_Quit();
    return 0;
}