# fsnotify
文件状态监控

## rock.fsnotify

- userdata = rock.fsnotify{name , path}
- name: 进程名
- path: 监控路径

#### 内部接口
- [userdata.start()]()
- [userdata.pipe(lua.writer)]()
- [userdata.on_err(lua.writer)]()

#### event
- [event.op string]()
- [event.name string]()
- [event.create bool]()
- [event.write bool]()
- [event.rename bool]()
- [event.chmod bool]()

```lua
    local ud = rock.fsnotify{
        name = "ff",
        path = "/var/log",
    }
    
    ud.pipe(function(ev)
        print(ev.op)
        print(ev.name) -- filename
        print(ev.write) -- true
    end)
    
    ud.on_err(function(err)
    end)

    ud.start()
```