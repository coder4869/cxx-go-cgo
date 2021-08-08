
#ifndef MACROS_H
#define MACROS_H

#ifndef API_EXPORT
#	if defined(__EMSCRIPTEN__)                  // 探测是否为 Emscripten 环境
#		include <emscripten/emscripten.h>       // 导出API必要头文件
#		if defined(__cplusplus)                 // 探测是否为 C++ 环境
#			define API_EXPORT(rettype) extern "C" rettype EMSCRIPTEN_KEEPALIVE  // EMSCRIPTEN_KEEPALIVE 是导出API必要声明
#		else
#			define API_EXPORT(rettype) rettype EMSCRIPTEN_KEEPALIVE
#		endif
#	else
#		if defined(__cplusplus)
#			define API_EXPORT(rettype) extern "C" rettype
#		else
#			define API_EXPORT(rettype) rettype
#		endif
#	endif
#endif

#endif // MACROS_H


