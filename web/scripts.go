// Code generated by "gen/scripts.go"; DO NOT EDIT.

package web

const(
// DefaultAppWorkersJS is the default template used in [MakeAppWorkerJS] to generate app-worker.js.
DefaultAppWorkerJS = "const cacheName = \"app-\" + \"{{.Version}}\";\r\nconst resourcesToCache = {{.ResourcesToCache}};\r\n\r\nself.addEventListener(\"install\", (event) => {\r\n  console.log(\"installing app worker {{.Version}}\");\r\n\r\n  event.waitUntil(\r\n    caches\r\n      .open(cacheName)\r\n      .then((cache) => {\r\n        return cache.addAll(resourcesToCache);\r\n      })\r\n      .then(() => {\r\n        self.skipWaiting();\r\n      })\r\n  );\r\n});\r\n\r\nself.addEventListener(\"activate\", (event) => {\r\n  event.waitUntil(\r\n    caches.keys().then((keyList) => {\r\n      return Promise.all(\r\n        keyList.map((key) => {\r\n          if (key !== cacheName) {\r\n            return caches.delete(key);\r\n          }\r\n        })\r\n      );\r\n    })\r\n  );\r\n  console.log(\"app worker {{.Version}} is activated\");\r\n});\r\n\r\nself.addEventListener(\"fetch\", (event) => {\r\n  event.respondWith(\r\n    caches.match(event.request).then((response) => {\r\n      return response || fetch(event.request);\r\n    })\r\n  );\r\n});\r\n\r\nself.addEventListener(\"push\", (event) => {\r\n  if (!event.data || !event.data.text()) {\r\n    return;\r\n  }\r\n\r\n  const notification = JSON.parse(event.data.text());\r\n  if (!notification) {\r\n    return;\r\n  }\r\n\r\n  const title = notification.title;\r\n  delete notification.title;\r\n\r\n  if (!notification.data) {\r\n    notification.data = {};\r\n  }\r\n  let actions = [];\r\n  for (let i in notification.actions) {\r\n    const action = notification.actions[i];\r\n\r\n    actions.push({\r\n      action: action.action,\r\n      path: action.path,\r\n    });\r\n\r\n    delete action.path;\r\n  }\r\n  notification.data.goapp = {\r\n    path: notification.path,\r\n    actions: actions,\r\n  };\r\n  delete notification.path;\r\n\r\n  event.waitUntil(self.registration.showNotification(title, notification));\r\n});\r\n\r\nself.addEventListener(\"notificationclick\", (event) => {\r\n  event.notification.close();\r\n\r\n  const notification = event.notification;\r\n  let path = notification.data.goapp.path;\r\n\r\n  for (let i in notification.data.goapp.actions) {\r\n    const action = notification.data.goapp.actions[i];\r\n    if (action.action === event.action) {\r\n      path = action.path;\r\n      break;\r\n    }\r\n  }\r\n\r\n  event.waitUntil(\r\n    clients\r\n      .matchAll({\r\n        type: \"window\",\r\n      })\r\n      .then((clientList) => {\r\n        for (var i = 0; i < clientList.length; i++) {\r\n          let client = clientList[i];\r\n          if (\"focus\" in client) {\r\n            client.focus();\r\n            client.postMessage({\r\n              goapp: {\r\n                type: \"notification\",\r\n                path: path,\r\n              },\r\n            });\r\n            return;\r\n          }\r\n        }\r\n\r\n        if (clients.openWindow) {\r\n          return clients.openWindow(path);\r\n        }\r\n      })\r\n  );\r\n});\r\n"

// WASMExecJSGoCurrent is the wasm_exec.js file for the current version of Go.
WASMExecJSGoCurrent = "// Copyright 2018 The Go Authors. All rights reserved.\n// Use of this source code is governed by a BSD-style\n// license that can be found in the LICENSE file.\n\n\"use strict\";\n\n(() => {\n\tconst enosys = () => {\n\t\tconst err = new Error(\"not implemented\");\n\t\terr.code = \"ENOSYS\";\n\t\treturn err;\n\t};\n\n\tif (!globalThis.fs) {\n\t\tlet outputBuf = \"\";\n\t\tglobalThis.fs = {\n\t\t\tconstants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1 }, // unused\n\t\t\twriteSync(fd, buf) {\n\t\t\t\toutputBuf += decoder.decode(buf);\n\t\t\t\tconst nl = outputBuf.lastIndexOf(\"\\n\");\n\t\t\t\tif (nl != -1) {\n\t\t\t\t\tconsole.log(outputBuf.substring(0, nl));\n\t\t\t\t\toutputBuf = outputBuf.substring(nl + 1);\n\t\t\t\t}\n\t\t\t\treturn buf.length;\n\t\t\t},\n\t\t\twrite(fd, buf, offset, length, position, callback) {\n\t\t\t\tif (offset !== 0 || length !== buf.length || position !== null) {\n\t\t\t\t\tcallback(enosys());\n\t\t\t\t\treturn;\n\t\t\t\t}\n\t\t\t\tconst n = this.writeSync(fd, buf);\n\t\t\t\tcallback(null, n);\n\t\t\t},\n\t\t\tchmod(path, mode, callback) { callback(enosys()); },\n\t\t\tchown(path, uid, gid, callback) { callback(enosys()); },\n\t\t\tclose(fd, callback) { callback(enosys()); },\n\t\t\tfchmod(fd, mode, callback) { callback(enosys()); },\n\t\t\tfchown(fd, uid, gid, callback) { callback(enosys()); },\n\t\t\tfstat(fd, callback) { callback(enosys()); },\n\t\t\tfsync(fd, callback) { callback(null); },\n\t\t\tftruncate(fd, length, callback) { callback(enosys()); },\n\t\t\tlchown(path, uid, gid, callback) { callback(enosys()); },\n\t\t\tlink(path, link, callback) { callback(enosys()); },\n\t\t\tlstat(path, callback) { callback(enosys()); },\n\t\t\tmkdir(path, perm, callback) { callback(enosys()); },\n\t\t\topen(path, flags, mode, callback) { callback(enosys()); },\n\t\t\tread(fd, buffer, offset, length, position, callback) { callback(enosys()); },\n\t\t\treaddir(path, callback) { callback(enosys()); },\n\t\t\treadlink(path, callback) { callback(enosys()); },\n\t\t\trename(from, to, callback) { callback(enosys()); },\n\t\t\trmdir(path, callback) { callback(enosys()); },\n\t\t\tstat(path, callback) { callback(enosys()); },\n\t\t\tsymlink(path, link, callback) { callback(enosys()); },\n\t\t\ttruncate(path, length, callback) { callback(enosys()); },\n\t\t\tunlink(path, callback) { callback(enosys()); },\n\t\t\tutimes(path, atime, mtime, callback) { callback(enosys()); },\n\t\t};\n\t}\n\n\tif (!globalThis.process) {\n\t\tglobalThis.process = {\n\t\t\tgetuid() { return -1; },\n\t\t\tgetgid() { return -1; },\n\t\t\tgeteuid() { return -1; },\n\t\t\tgetegid() { return -1; },\n\t\t\tgetgroups() { throw enosys(); },\n\t\t\tpid: -1,\n\t\t\tppid: -1,\n\t\t\tumask() { throw enosys(); },\n\t\t\tcwd() { throw enosys(); },\n\t\t\tchdir() { throw enosys(); },\n\t\t}\n\t}\n\n\tif (!globalThis.crypto) {\n\t\tthrow new Error(\"globalThis.crypto is not available, polyfill required (crypto.getRandomValues only)\");\n\t}\n\n\tif (!globalThis.performance) {\n\t\tthrow new Error(\"globalThis.performance is not available, polyfill required (performance.now only)\");\n\t}\n\n\tif (!globalThis.TextEncoder) {\n\t\tthrow new Error(\"globalThis.TextEncoder is not available, polyfill required\");\n\t}\n\n\tif (!globalThis.TextDecoder) {\n\t\tthrow new Error(\"globalThis.TextDecoder is not available, polyfill required\");\n\t}\n\n\tconst encoder = new TextEncoder(\"utf-8\");\n\tconst decoder = new TextDecoder(\"utf-8\");\n\n\tglobalThis.Go = class {\n\t\tconstructor() {\n\t\t\tthis.argv = [\"js\"];\n\t\t\tthis.env = {};\n\t\t\tthis.exit = (code) => {\n\t\t\t\tif (code !== 0) {\n\t\t\t\t\tconsole.warn(\"exit code:\", code);\n\t\t\t\t}\n\t\t\t};\n\t\t\tthis._exitPromise = new Promise((resolve) => {\n\t\t\t\tthis._resolveExitPromise = resolve;\n\t\t\t});\n\t\t\tthis._pendingEvent = null;\n\t\t\tthis._scheduledTimeouts = new Map();\n\t\t\tthis._nextCallbackTimeoutID = 1;\n\n\t\t\tconst setInt64 = (addr, v) => {\n\t\t\t\tthis.mem.setUint32(addr + 0, v, true);\n\t\t\t\tthis.mem.setUint32(addr + 4, Math.floor(v / 4294967296), true);\n\t\t\t}\n\n\t\t\tconst setInt32 = (addr, v) => {\n\t\t\t\tthis.mem.setUint32(addr + 0, v, true);\n\t\t\t}\n\n\t\t\tconst getInt64 = (addr) => {\n\t\t\t\tconst low = this.mem.getUint32(addr + 0, true);\n\t\t\t\tconst high = this.mem.getInt32(addr + 4, true);\n\t\t\t\treturn low + high * 4294967296;\n\t\t\t}\n\n\t\t\tconst loadValue = (addr) => {\n\t\t\t\tconst f = this.mem.getFloat64(addr, true);\n\t\t\t\tif (f === 0) {\n\t\t\t\t\treturn undefined;\n\t\t\t\t}\n\t\t\t\tif (!isNaN(f)) {\n\t\t\t\t\treturn f;\n\t\t\t\t}\n\n\t\t\t\tconst id = this.mem.getUint32(addr, true);\n\t\t\t\treturn this._values[id];\n\t\t\t}\n\n\t\t\tconst storeValue = (addr, v) => {\n\t\t\t\tconst nanHead = 0x7FF80000;\n\n\t\t\t\tif (typeof v === \"number\" && v !== 0) {\n\t\t\t\t\tif (isNaN(v)) {\n\t\t\t\t\t\tthis.mem.setUint32(addr + 4, nanHead, true);\n\t\t\t\t\t\tthis.mem.setUint32(addr, 0, true);\n\t\t\t\t\t\treturn;\n\t\t\t\t\t}\n\t\t\t\t\tthis.mem.setFloat64(addr, v, true);\n\t\t\t\t\treturn;\n\t\t\t\t}\n\n\t\t\t\tif (v === undefined) {\n\t\t\t\t\tthis.mem.setFloat64(addr, 0, true);\n\t\t\t\t\treturn;\n\t\t\t\t}\n\n\t\t\t\tlet id = this._ids.get(v);\n\t\t\t\tif (id === undefined) {\n\t\t\t\t\tid = this._idPool.pop();\n\t\t\t\t\tif (id === undefined) {\n\t\t\t\t\t\tid = this._values.length;\n\t\t\t\t\t}\n\t\t\t\t\tthis._values[id] = v;\n\t\t\t\t\tthis._goRefCounts[id] = 0;\n\t\t\t\t\tthis._ids.set(v, id);\n\t\t\t\t}\n\t\t\t\tthis._goRefCounts[id]++;\n\t\t\t\tlet typeFlag = 0;\n\t\t\t\tswitch (typeof v) {\n\t\t\t\t\tcase \"object\":\n\t\t\t\t\t\tif (v !== null) {\n\t\t\t\t\t\t\ttypeFlag = 1;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tbreak;\n\t\t\t\t\tcase \"string\":\n\t\t\t\t\t\ttypeFlag = 2;\n\t\t\t\t\t\tbreak;\n\t\t\t\t\tcase \"symbol\":\n\t\t\t\t\t\ttypeFlag = 3;\n\t\t\t\t\t\tbreak;\n\t\t\t\t\tcase \"function\":\n\t\t\t\t\t\ttypeFlag = 4;\n\t\t\t\t\t\tbreak;\n\t\t\t\t}\n\t\t\t\tthis.mem.setUint32(addr + 4, nanHead | typeFlag, true);\n\t\t\t\tthis.mem.setUint32(addr, id, true);\n\t\t\t}\n\n\t\t\tconst loadSlice = (addr) => {\n\t\t\t\tconst array = getInt64(addr + 0);\n\t\t\t\tconst len = getInt64(addr + 8);\n\t\t\t\treturn new Uint8Array(this._inst.exports.mem.buffer, array, len);\n\t\t\t}\n\n\t\t\tconst loadSliceOfValues = (addr) => {\n\t\t\t\tconst array = getInt64(addr + 0);\n\t\t\t\tconst len = getInt64(addr + 8);\n\t\t\t\tconst a = new Array(len);\n\t\t\t\tfor (let i = 0; i < len; i++) {\n\t\t\t\t\ta[i] = loadValue(array + i * 8);\n\t\t\t\t}\n\t\t\t\treturn a;\n\t\t\t}\n\n\t\t\tconst loadString = (addr) => {\n\t\t\t\tconst saddr = getInt64(addr + 0);\n\t\t\t\tconst len = getInt64(addr + 8);\n\t\t\t\treturn decoder.decode(new DataView(this._inst.exports.mem.buffer, saddr, len));\n\t\t\t}\n\n\t\t\tconst timeOrigin = Date.now() - performance.now();\n\t\t\tthis.importObject = {\n\t\t\t\t_gotest: {\n\t\t\t\t\tadd: (a, b) => a + b,\n\t\t\t\t},\n\t\t\t\tgojs: {\n\t\t\t\t\t// Go's SP does not change as long as no Go code is running. Some operations (e.g. calls, getters and setters)\n\t\t\t\t\t// may synchronously trigger a Go event handler. This makes Go code get executed in the middle of the imported\n\t\t\t\t\t// function. A goroutine can switch to a new stack if the current stack is too small (see morestack function).\n\t\t\t\t\t// This changes the SP, thus we have to update the SP used by the imported function.\n\n\t\t\t\t\t// func wasmExit(code int32)\n\t\t\t\t\t\"runtime.wasmExit\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst code = this.mem.getInt32(sp + 8, true);\n\t\t\t\t\t\tthis.exited = true;\n\t\t\t\t\t\tdelete this._inst;\n\t\t\t\t\t\tdelete this._values;\n\t\t\t\t\t\tdelete this._goRefCounts;\n\t\t\t\t\t\tdelete this._ids;\n\t\t\t\t\t\tdelete this._idPool;\n\t\t\t\t\t\tthis.exit(code);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func wasmWrite(fd uintptr, p unsafe.Pointer, n int32)\n\t\t\t\t\t\"runtime.wasmWrite\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst fd = getInt64(sp + 8);\n\t\t\t\t\t\tconst p = getInt64(sp + 16);\n\t\t\t\t\t\tconst n = this.mem.getInt32(sp + 24, true);\n\t\t\t\t\t\tfs.writeSync(fd, new Uint8Array(this._inst.exports.mem.buffer, p, n));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func resetMemoryDataView()\n\t\t\t\t\t\"runtime.resetMemoryDataView\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tthis.mem = new DataView(this._inst.exports.mem.buffer);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func nanotime1() int64\n\t\t\t\t\t\"runtime.nanotime1\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tsetInt64(sp + 8, (timeOrigin + performance.now()) * 1000000);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func walltime() (sec int64, nsec int32)\n\t\t\t\t\t\"runtime.walltime\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst msec = (new Date).getTime();\n\t\t\t\t\t\tsetInt64(sp + 8, msec / 1000);\n\t\t\t\t\t\tthis.mem.setInt32(sp + 16, (msec % 1000) * 1000000, true);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func scheduleTimeoutEvent(delay int64) int32\n\t\t\t\t\t\"runtime.scheduleTimeoutEvent\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst id = this._nextCallbackTimeoutID;\n\t\t\t\t\t\tthis._nextCallbackTimeoutID++;\n\t\t\t\t\t\tthis._scheduledTimeouts.set(id, setTimeout(\n\t\t\t\t\t\t\t() => {\n\t\t\t\t\t\t\t\tthis._resume();\n\t\t\t\t\t\t\t\twhile (this._scheduledTimeouts.has(id)) {\n\t\t\t\t\t\t\t\t\t// for some reason Go failed to register the timeout event, log and try again\n\t\t\t\t\t\t\t\t\t// (temporary workaround for https://github.com/golang/go/issues/28975)\n\t\t\t\t\t\t\t\t\tconsole.warn(\"scheduleTimeoutEvent: missed timeout event\");\n\t\t\t\t\t\t\t\t\tthis._resume();\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\tgetInt64(sp + 8),\n\t\t\t\t\t\t));\n\t\t\t\t\t\tthis.mem.setInt32(sp + 16, id, true);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func clearTimeoutEvent(id int32)\n\t\t\t\t\t\"runtime.clearTimeoutEvent\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst id = this.mem.getInt32(sp + 8, true);\n\t\t\t\t\t\tclearTimeout(this._scheduledTimeouts.get(id));\n\t\t\t\t\t\tthis._scheduledTimeouts.delete(id);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func getRandomData(r []byte)\n\t\t\t\t\t\"runtime.getRandomData\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tcrypto.getRandomValues(loadSlice(sp + 8));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func finalizeRef(v ref)\n\t\t\t\t\t\"syscall/js.finalizeRef\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst id = this.mem.getUint32(sp + 8, true);\n\t\t\t\t\t\tthis._goRefCounts[id]--;\n\t\t\t\t\t\tif (this._goRefCounts[id] === 0) {\n\t\t\t\t\t\t\tconst v = this._values[id];\n\t\t\t\t\t\t\tthis._values[id] = null;\n\t\t\t\t\t\t\tthis._ids.delete(v);\n\t\t\t\t\t\t\tthis._idPool.push(id);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\t// func stringVal(value string) ref\n\t\t\t\t\t\"syscall/js.stringVal\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tstoreValue(sp + 24, loadString(sp + 8));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueGet(v ref, p string) ref\n\t\t\t\t\t\"syscall/js.valueGet\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst result = Reflect.get(loadValue(sp + 8), loadString(sp + 16));\n\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\tstoreValue(sp + 32, result);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueSet(v ref, p string, x ref)\n\t\t\t\t\t\"syscall/js.valueSet\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tReflect.set(loadValue(sp + 8), loadString(sp + 16), loadValue(sp + 32));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueDelete(v ref, p string)\n\t\t\t\t\t\"syscall/js.valueDelete\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tReflect.deleteProperty(loadValue(sp + 8), loadString(sp + 16));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueIndex(v ref, i int) ref\n\t\t\t\t\t\"syscall/js.valueIndex\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tstoreValue(sp + 24, Reflect.get(loadValue(sp + 8), getInt64(sp + 16)));\n\t\t\t\t\t},\n\n\t\t\t\t\t// valueSetIndex(v ref, i int, x ref)\n\t\t\t\t\t\"syscall/js.valueSetIndex\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tReflect.set(loadValue(sp + 8), getInt64(sp + 16), loadValue(sp + 24));\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueCall(v ref, m string, args []ref) (ref, bool)\n\t\t\t\t\t\"syscall/js.valueCall\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tconst v = loadValue(sp + 8);\n\t\t\t\t\t\t\tconst m = Reflect.get(v, loadString(sp + 16));\n\t\t\t\t\t\t\tconst args = loadSliceOfValues(sp + 32);\n\t\t\t\t\t\t\tconst result = Reflect.apply(m, v, args);\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 56, result);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 64, 1);\n\t\t\t\t\t\t} catch (err) {\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 56, err);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 64, 0);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueInvoke(v ref, args []ref) (ref, bool)\n\t\t\t\t\t\"syscall/js.valueInvoke\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tconst v = loadValue(sp + 8);\n\t\t\t\t\t\t\tconst args = loadSliceOfValues(sp + 16);\n\t\t\t\t\t\t\tconst result = Reflect.apply(v, undefined, args);\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 40, result);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 1);\n\t\t\t\t\t\t} catch (err) {\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 40, err);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 0);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueNew(v ref, args []ref) (ref, bool)\n\t\t\t\t\t\"syscall/js.valueNew\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tconst v = loadValue(sp + 8);\n\t\t\t\t\t\t\tconst args = loadSliceOfValues(sp + 16);\n\t\t\t\t\t\t\tconst result = Reflect.construct(v, args);\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 40, result);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 1);\n\t\t\t\t\t\t} catch (err) {\n\t\t\t\t\t\t\tsp = this._inst.exports.getsp() >>> 0; // see comment above\n\t\t\t\t\t\t\tstoreValue(sp + 40, err);\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 0);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueLength(v ref) int\n\t\t\t\t\t\"syscall/js.valueLength\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tsetInt64(sp + 16, parseInt(loadValue(sp + 8).length));\n\t\t\t\t\t},\n\n\t\t\t\t\t// valuePrepareString(v ref) (ref, int)\n\t\t\t\t\t\"syscall/js.valuePrepareString\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst str = encoder.encode(String(loadValue(sp + 8)));\n\t\t\t\t\t\tstoreValue(sp + 16, str);\n\t\t\t\t\t\tsetInt64(sp + 24, str.length);\n\t\t\t\t\t},\n\n\t\t\t\t\t// valueLoadString(v ref, b []byte)\n\t\t\t\t\t\"syscall/js.valueLoadString\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst str = loadValue(sp + 8);\n\t\t\t\t\t\tloadSlice(sp + 16).set(str);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func valueInstanceOf(v ref, t ref) bool\n\t\t\t\t\t\"syscall/js.valueInstanceOf\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tthis.mem.setUint8(sp + 24, (loadValue(sp + 8) instanceof loadValue(sp + 16)) ? 1 : 0);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func copyBytesToGo(dst []byte, src ref) (int, bool)\n\t\t\t\t\t\"syscall/js.copyBytesToGo\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst dst = loadSlice(sp + 8);\n\t\t\t\t\t\tconst src = loadValue(sp + 32);\n\t\t\t\t\t\tif (!(src instanceof Uint8Array || src instanceof Uint8ClampedArray)) {\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 0);\n\t\t\t\t\t\t\treturn;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tconst toCopy = src.subarray(0, dst.length);\n\t\t\t\t\t\tdst.set(toCopy);\n\t\t\t\t\t\tsetInt64(sp + 40, toCopy.length);\n\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 1);\n\t\t\t\t\t},\n\n\t\t\t\t\t// func copyBytesToJS(dst ref, src []byte) (int, bool)\n\t\t\t\t\t\"syscall/js.copyBytesToJS\": (sp) => {\n\t\t\t\t\t\tsp >>>= 0;\n\t\t\t\t\t\tconst dst = loadValue(sp + 8);\n\t\t\t\t\t\tconst src = loadSlice(sp + 16);\n\t\t\t\t\t\tif (!(dst instanceof Uint8Array || dst instanceof Uint8ClampedArray)) {\n\t\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 0);\n\t\t\t\t\t\t\treturn;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tconst toCopy = src.subarray(0, dst.length);\n\t\t\t\t\t\tdst.set(toCopy);\n\t\t\t\t\t\tsetInt64(sp + 40, toCopy.length);\n\t\t\t\t\t\tthis.mem.setUint8(sp + 48, 1);\n\t\t\t\t\t},\n\n\t\t\t\t\t\"debug\": (value) => {\n\t\t\t\t\t\tconsole.log(value);\n\t\t\t\t\t},\n\t\t\t\t}\n\t\t\t};\n\t\t}\n\n\t\tasync run(instance) {\n\t\t\tif (!(instance instanceof WebAssembly.Instance)) {\n\t\t\t\tthrow new Error(\"Go.run: WebAssembly.Instance expected\");\n\t\t\t}\n\t\t\tthis._inst = instance;\n\t\t\tthis.mem = new DataView(this._inst.exports.mem.buffer);\n\t\t\tthis._values = [ // JS values that Go currently has references to, indexed by reference id\n\t\t\t\tNaN,\n\t\t\t\t0,\n\t\t\t\tnull,\n\t\t\t\ttrue,\n\t\t\t\tfalse,\n\t\t\t\tglobalThis,\n\t\t\t\tthis,\n\t\t\t];\n\t\t\tthis._goRefCounts = new Array(this._values.length).fill(Infinity); // number of references that Go has to a JS value, indexed by reference id\n\t\t\tthis._ids = new Map([ // mapping from JS values to reference ids\n\t\t\t\t[0, 1],\n\t\t\t\t[null, 2],\n\t\t\t\t[true, 3],\n\t\t\t\t[false, 4],\n\t\t\t\t[globalThis, 5],\n\t\t\t\t[this, 6],\n\t\t\t]);\n\t\t\tthis._idPool = [];   // unused ids that have been garbage collected\n\t\t\tthis.exited = false; // whether the Go program has exited\n\n\t\t\t// Pass command line arguments and environment variables to WebAssembly by writing them to the linear memory.\n\t\t\tlet offset = 4096;\n\n\t\t\tconst strPtr = (str) => {\n\t\t\t\tconst ptr = offset;\n\t\t\t\tconst bytes = encoder.encode(str + \"\\0\");\n\t\t\t\tnew Uint8Array(this.mem.buffer, offset, bytes.length).set(bytes);\n\t\t\t\toffset += bytes.length;\n\t\t\t\tif (offset % 8 !== 0) {\n\t\t\t\t\toffset += 8 - (offset % 8);\n\t\t\t\t}\n\t\t\t\treturn ptr;\n\t\t\t};\n\n\t\t\tconst argc = this.argv.length;\n\n\t\t\tconst argvPtrs = [];\n\t\t\tthis.argv.forEach((arg) => {\n\t\t\t\targvPtrs.push(strPtr(arg));\n\t\t\t});\n\t\t\targvPtrs.push(0);\n\n\t\t\tconst keys = Object.keys(this.env).sort();\n\t\t\tkeys.forEach((key) => {\n\t\t\t\targvPtrs.push(strPtr(`${key}=${this.env[key]}`));\n\t\t\t});\n\t\t\targvPtrs.push(0);\n\n\t\t\tconst argv = offset;\n\t\t\targvPtrs.forEach((ptr) => {\n\t\t\t\tthis.mem.setUint32(offset, ptr, true);\n\t\t\t\tthis.mem.setUint32(offset + 4, 0, true);\n\t\t\t\toffset += 8;\n\t\t\t});\n\n\t\t\t// The linker guarantees global data starts from at least wasmMinDataAddr.\n\t\t\t// Keep in sync with cmd/link/internal/ld/data.go:wasmMinDataAddr.\n\t\t\tconst wasmMinDataAddr = 4096 + 8192;\n\t\t\tif (offset >= wasmMinDataAddr) {\n\t\t\t\tthrow new Error(\"total length of command line and environment variables exceeds limit\");\n\t\t\t}\n\n\t\t\tthis._inst.exports.run(argc, argv);\n\t\t\tif (this.exited) {\n\t\t\t\tthis._resolveExitPromise();\n\t\t\t}\n\t\t\tawait this._exitPromise;\n\t\t}\n\n\t\t_resume() {\n\t\t\tif (this.exited) {\n\t\t\t\tthrow new Error(\"Go program has already exited\");\n\t\t\t}\n\t\t\tthis._inst.exports.resume();\n\t\t\tif (this.exited) {\n\t\t\t\tthis._resolveExitPromise();\n\t\t\t}\n\t\t}\n\n\t\t_makeFuncWrapper(id) {\n\t\t\tconst go = this;\n\t\t\treturn function () {\n\t\t\t\tconst event = { id: id, this: this, args: arguments };\n\t\t\t\tgo._pendingEvent = event;\n\t\t\t\tgo._resume();\n\t\t\t\treturn event.result;\n\t\t\t};\n\t\t}\n\t}\n})();\n"

// AppJS is the string used for [AppJSTmpl].
AppJS = "// -----------------------------------------------------------------------------\r\n// go-app\r\n// -----------------------------------------------------------------------------\r\nvar goappNav = function () {};\r\nvar goappOnUpdate = function () {};\r\nvar goappOnAppInstallChange = function () {};\r\n\r\nconst goappEnv = {{.Env}};\r\nconst goappLoadingLabel = \"{{.LoadingLabel}}\";\r\nconst goappWasmContentLengthHeader = \"{{.WasmContentLengthHeader}}\";\r\n\r\nlet goappServiceWorkerRegistration;\r\nlet deferredPrompt = null;\r\n\r\ngoappInitServiceWorker();\r\ngoappWatchForUpdate();\r\ngoappWatchForInstallable();\r\ngoappInitWebAssembly();\r\n\r\n// -----------------------------------------------------------------------------\r\n// Service Worker\r\n// -----------------------------------------------------------------------------\r\nasync function goappInitServiceWorker() {\r\n  if (\"serviceWorker\" in navigator) {\r\n    try {\r\n      const registration = await navigator.serviceWorker.register(\r\n        \"{{.WorkerJS}}\"\r\n      );\r\n\r\n      goappServiceWorkerRegistration = registration;\r\n      goappSetupNotifyUpdate(registration);\r\n      goappSetupAutoUpdate(registration);\r\n      goappSetupPushNotification();\r\n    } catch (err) {\r\n      console.error(\"goapp service worker registration failed\", err);\r\n    }\r\n  }\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Update\r\n// -----------------------------------------------------------------------------\r\nfunction goappWatchForUpdate() {\r\n  window.addEventListener(\"beforeinstallprompt\", (e) => {\r\n    e.preventDefault();\r\n    deferredPrompt = e;\r\n    goappOnAppInstallChange();\r\n  });\r\n}\r\n\r\nfunction goappSetupNotifyUpdate(registration) {\r\n  registration.addEventListener(\"updatefound\", (event) => {\r\n    const newSW = registration.installing;\r\n    newSW.addEventListener(\"statechange\", (event) => {\r\n      if (!navigator.serviceWorker.controller) {\r\n        return;\r\n      }\r\n      if (newSW.state != \"installed\") {\r\n        return;\r\n      }\r\n      goappOnUpdate();\r\n    });\r\n  });\r\n}\r\n\r\nfunction goappSetupAutoUpdate(registration) {\r\n  const autoUpdateInterval = \"{{.AutoUpdateInterval}}\";\r\n  if (autoUpdateInterval == 0) {\r\n    return;\r\n  }\r\n\r\n  window.setInterval(() => {\r\n    registration.update();\r\n  }, autoUpdateInterval);\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Install\r\n// -----------------------------------------------------------------------------\r\nfunction goappWatchForInstallable() {\r\n  window.addEventListener(\"appinstalled\", () => {\r\n    deferredPrompt = null;\r\n    goappOnAppInstallChange();\r\n  });\r\n}\r\n\r\nfunction goappIsAppInstallable() {\r\n  return !goappIsAppInstalled() && deferredPrompt != null;\r\n}\r\n\r\nfunction goappIsAppInstalled() {\r\n  const isStandalone = window.matchMedia(\"(display-mode: standalone)\").matches;\r\n  return isStandalone || navigator.standalone;\r\n}\r\n\r\nasync function goappShowInstallPrompt() {\r\n  deferredPrompt.prompt();\r\n  await deferredPrompt.userChoice;\r\n  deferredPrompt = null;\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Environment\r\n// -----------------------------------------------------------------------------\r\nfunction goappGetenv(k) {\r\n  return goappEnv[k];\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Notifications\r\n// -----------------------------------------------------------------------------\r\nfunction goappSetupPushNotification() {\r\n  navigator.serviceWorker.addEventListener(\"message\", (event) => {\r\n    const msg = event.data.goapp;\r\n    if (!msg) {\r\n      return;\r\n    }\r\n\r\n    if (msg.type !== \"notification\") {\r\n      return;\r\n    }\r\n\r\n    goappNav(msg.path);\r\n  });\r\n}\r\n\r\nasync function goappSubscribePushNotifications(vapIDpublicKey) {\r\n  try {\r\n    const subscription =\r\n      await goappServiceWorkerRegistration.pushManager.subscribe({\r\n        userVisibleOnly: true,\r\n        applicationServerKey: vapIDpublicKey,\r\n      });\r\n    return JSON.stringify(subscription);\r\n  } catch (err) {\r\n    console.error(err);\r\n    return \"\";\r\n  }\r\n}\r\n\r\nfunction goappNewNotification(jsonNotification) {\r\n  let notification = JSON.parse(jsonNotification);\r\n\r\n  const title = notification.title;\r\n  delete notification.title;\r\n\r\n  let path = notification.path;\r\n  if (!path) {\r\n    path = \"/\";\r\n  }\r\n\r\n  const webNotification = new Notification(title, notification);\r\n\r\n  webNotification.onclick = () => {\r\n    goappNav(path);\r\n    webNotification.close();\r\n  };\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Keep Clean Body\r\n// -----------------------------------------------------------------------------\r\nfunction goappKeepBodyClean() {\r\n  const body = document.body;\r\n  const bodyChildrenCount = body.children.length;\r\n\r\n  const mutationObserver = new MutationObserver(function (mutationList) {\r\n    mutationList.forEach((mutation) => {\r\n      switch (mutation.type) {\r\n        case \"childList\":\r\n          while (body.children.length > bodyChildrenCount) {\r\n            body.removeChild(body.lastChild);\r\n          }\r\n          break;\r\n      }\r\n    });\r\n  });\r\n\r\n  mutationObserver.observe(document.body, {\r\n    childList: true,\r\n  });\r\n\r\n  return () => mutationObserver.disconnect();\r\n}\r\n\r\n// -----------------------------------------------------------------------------\r\n// Web Assembly\r\n// -----------------------------------------------------------------------------\r\nasync function goappInitWebAssembly() {\r\n  const loader = document.getElementById(\"app-wasm-loader\");\r\n\r\n  if (!goappCanLoadWebAssembly()) {\r\n    loader.remove();\r\n    return;\r\n  }\r\n\r\n  let instantiateStreaming = WebAssembly.instantiateStreaming;\r\n  if (!instantiateStreaming) {\r\n    instantiateStreaming = async (resp, importObject) => {\r\n      const source = await (await resp).arrayBuffer();\r\n      return await WebAssembly.instantiate(source, importObject);\r\n    };\r\n  }\r\n\r\n  const loaderIcon = document.getElementById(\"app-wasm-loader-icon\");\r\n  const loaderLabel = document.getElementById(\"app-wasm-loader-label\");\r\n\r\n  try {\r\n    const showProgress = (progress) => {\r\n      loaderLabel.innerText = goappLoadingLabel.replace(\"{progress}\", progress);\r\n    };\r\n    showProgress(0);\r\n\r\n    const go = new Go();\r\n    const wasm = await instantiateStreaming(\r\n      fetchWithProgress(\"{{.Wasm}}\", showProgress),\r\n      go.importObject\r\n    );\r\n\r\n    go.run(wasm.instance);\r\n    loader.remove();\r\n  } catch (err) {\r\n    loaderIcon.className = \"goapp-logo\";\r\n    loaderLabel.innerText = err;\r\n    console.error(\"loading wasm failed: \", err);\r\n  }\r\n}\r\n\r\nfunction goappCanLoadWebAssembly() {\r\n  if (\r\n    /bot|googlebot|crawler|spider|robot|crawling/i.test(navigator.userAgent)\r\n  ) {\r\n    return false;\r\n  }\r\n\r\n  const urlParams = new URLSearchParams(window.location.search);\r\n  return urlParams.get(\"wasm\") !== \"false\";\r\n}\r\n\r\nasync function fetchWithProgress(url, progess) {\r\n  const response = await fetch(url);\r\n\r\n  let contentLength;\r\n  try {\r\n    contentLength = response.headers.get(goappWasmContentLengthHeader);\r\n  } catch {}\r\n  if (!goappWasmContentLengthHeader || !contentLength) {\r\n    contentLength = response.headers.get(\"Content-Length\");\r\n  }\r\n\r\n  const total = parseInt(contentLength, 10);\r\n  let loaded = 0;\r\n\r\n  const progressHandler = function (loaded, total) {\r\n    progess(Math.round((loaded * 100) / total));\r\n  };\r\n\r\n  var res = new Response(\r\n    new ReadableStream(\r\n      {\r\n        async start(controller) {\r\n          var reader = response.body.getReader();\r\n          for (;;) {\r\n            var { done, value } = await reader.read();\r\n\r\n            if (done) {\r\n              progressHandler(total, total);\r\n              break;\r\n            }\r\n\r\n            loaded += value.byteLength;\r\n            progressHandler(loaded, total);\r\n            controller.enqueue(value);\r\n          }\r\n          controller.close();\r\n        },\r\n      },\r\n      {\r\n        status: response.status,\r\n        statusText: response.statusText,\r\n      }\r\n    )\r\n  );\r\n\r\n  for (var pair of response.headers.entries()) {\r\n    res.headers.set(pair[0], pair[1]);\r\n  }\r\n\r\n  return res;\r\n}\r\n"

// ManifestJSON is the string used for [ManifestJSONTmpl].
ManifestJSON = "{\r\n  \"short_name\": \"{{.ShortName}}\",\r\n  \"name\": \"{{.Name}}\",\r\n  \"description\": \"{{.Description}}\",\r\n  \"icons\": [\r\n    {\r\n      \"src\": \"{{.SVGIcon}}\",\r\n      \"type\": \"image/svg+xml\",\r\n      \"sizes\": \"any\"\r\n    },\r\n    {\r\n      \"src\": \"{{.LargeIcon}}\",\r\n      \"type\": \"image/png\",\r\n      \"sizes\": \"512x512\"\r\n    },\r\n    {\r\n      \"src\": \"{{.DefaultIcon}}\",\r\n      \"type\": \"image/png\",\r\n      \"sizes\": \"192x192\"\r\n    }\r\n  ],\r\n  \"scope\": \"{{.Scope}}\",\r\n  \"start_url\": \"{{.StartURL}}\",\r\n  \"background_color\": \"{{.BackgroundColor}}\",\r\n  \"theme_color\": \"{{.ThemeColor}}\",\r\n  \"display\": \"standalone\"\r\n}"

// AppCSS is the string used for app.css.
AppCSS = "/*------------------------------------------------------------------------------\r\n  Loader\r\n------------------------------------------------------------------------------*/\r\n.goapp-app-info {\r\n  position: fixed;\r\n  top: 0;\r\n  left: 0;\r\n  z-index: 1000;\r\n  width: 100vw;\r\n  height: 100vh;\r\n  overflow: hidden;\r\n\r\n  display: flex;\r\n  flex-direction: column;\r\n  justify-content: center;\r\n  align-items: center;\r\n\r\n  font-family: -apple-system, BlinkMacSystemFont, \"Segoe UI\", Roboto, Oxygen,\r\n    Ubuntu, Cantarell, \"Open Sans\", \"Helvetica Neue\", sans-serif;\r\n  font-size: 13px;\r\n  font-weight: 400;\r\n  color: white;\r\n  background-color: #2d2c2c;\r\n}\r\n\r\n@media (prefers-color-scheme: light) {\r\n  .goapp-app-info {\r\n    color: black;\r\n    background-color: #f6f6f6;\r\n  }\r\n}\r\n\r\n.goapp-logo {\r\n  max-width: 100px;\r\n  max-height: 100px;\r\n  user-select: none;\r\n  -moz-user-select: none;\r\n  -webkit-user-drag: none;\r\n  -webkit-user-select: none;\r\n  -ms-user-select: none;\r\n}\r\n\r\n.goapp-label {\r\n  margin-top: 12px;\r\n  font-size: 21px;\r\n  font-weight: 100;\r\n  letter-spacing: 1px;\r\n  max-width: 480px;\r\n  text-align: center;\r\n}\r\n\r\n.goapp-spin {\r\n  animation: goapp-spin-frames 1.21s infinite linear;\r\n}\r\n\r\n@keyframes goapp-spin-frames {\r\n  from {\r\n    transform: rotate(0deg);\r\n  }\r\n\r\n  to {\r\n    transform: rotate(360deg);\r\n  }\r\n}\r\n\r\n/*------------------------------------------------------------------------------\r\n  Not found\r\n------------------------------------------------------------------------------*/\r\n.goapp-notfound-title {\r\n  display: flex;\r\n  justify-content: center;\r\n  align-items: center;\r\n  font-size: 65pt;\r\n  font-weight: 100;\r\n}\r\n"

// IndexHTML is the string used for [IndexHTMLTmpl].
IndexHTML = "<!DOCTYPE html>\r\n<html lang=\"en\">\r\n\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"author\" content=\"{{.Author}}\">\r\n    <meta name=\"description\" content=\"{{.Desc}}\">\r\n    <meta name=\"keywords\" content=\"{{.Keywords}}\">\r\n    <meta name=\"theme-color\" content=\"{{.ThemeColor}}\">\r\n    <meta name=\"viewport\"\r\n        content=\"width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover\">\r\n    <meta property=\"og:url\" content=\"http://127.0.0.1:60452/\">\r\n    <meta property=\"og:title\" content=\"{{.Title}}\">\r\n    <meta property=\"og:description\" content=\"{{.Desc}}\">\r\n    <meta property=\"og:type\" content=\"website\">\r\n    <meta property=\"og:image\" content=\"{{.Image}}\">\r\n    <title>{{.Title}}</title>\r\n    <link type=\"text/css\" rel=\"preload\" href=\"app.css\" as=\"style\">\r\n    <link rel=\"icon\" href=\"https://raw.githubusercontent.com/maxence-charriere/go-app/master/docs/web/icon.svg\">\r\n    <link rel=\"apple-touch-icon\" href=\"/web/icon.png\">\r\n    <link rel=\"manifest\" href=\"manifest.webmanifest\">\r\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"app.css\">\r\n    <script defer src=\"wasm_exec.js\"></script>\r\n    <script defer src=\"app.js\"></script>\r\n</head>\r\n\r\n<body>\r\n    <aside id=\"app-wasm-loader\" class=\"goapp-app-info\">\r\n        <img class=\"goapp-logo goapp-spin\" src=\"/web/icon.png\" id=\"app-wasm-loader-icon\">\r\n        <p id=\"app-wasm-loader-label\" class=\"goapp-label\">Loading {{.Title}}...</p>\r\n    </aside>\r\n</body>\r\n\r\n</html>"

)
