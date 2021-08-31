#ifndef VEC_STRUCT_H
#define VEC_STRUCT_H

// https://blog.csdn.net/guniwi/article/details/109806439

// define struct with `struct`
struct Vec2i {
    int x;
    int y;
};


// define struct with `typedef struct`
typedef struct Vec2fTag {
    float x;
    float y;
    // Vec2f(float _x, float _y);
} Vec2f;

#endif // VEC_STRUCT_H