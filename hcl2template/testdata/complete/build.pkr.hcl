
// starts resources to provision them.
build {
    sources = [ 
        "source.virtualbox-iso.ubuntu-1204",
    ]

    provisioner "shell" {
        not_squashed = var.foo
        string   = "string"
        int      = "${41 + 1}"
        int64    = "${42 + 1}"
        bool     = "true"
        trilean  = true
        duration = "${9 + 1}s"
        map_string_string {
            a = "b"
            c = "d"
        }
        slice_string = var.availability_zone_names

        nested {
            string   = "string"
            int      = 42
            int64    = 43
            bool     = true
            trilean  = true
            duration = "10s"
            map_string_string {
                a = "b"
                c = "d"
            }
            slice_string = var.availability_zone_names
        }

        nested_slice {
        }
    }

    provisioner "file" {
        not_squashed = "${var.foo}"
        string   = "string"
        int      = 42
        int64    = 43
        bool     = true
        trilean  = true
        duration = "10s"
        map_string_string {
            a = "b"
            c = "d"
        }
        slice_string = [
            "a",
            "b",
            "c",
        ]

        nested {
            string   = "string"
            int      = 42
            int64    = 43
            bool     = true
            trilean  = true
            duration = "10s"
            map_string_string {
                a = "b"
                c = "d"
            }
            slice_string = [
                "a",
                "b",
                "c",
            ]
        }

        nested_slice {
        }
    }

    post-processor "amazon-import" { 
        string   = "string"
        int      = 42
        int64    = 43
        bool     = true
        trilean  = true
        duration = "10s"
        map_string_string {
            a = "b"
            c = "d"
        }
        slice_string = [
            "a",
            "b",
            "c",
        ]

        nested {
            string   = "string"
            int      = 42
            int64    = 43
            bool     = true
            trilean  = true
            duration = "10s"
            map_string_string {
                a = "b"
                c = "d"
            }
            slice_string = [
                "a",
                "b",
                "c",
            ]
        }

        nested_slice {
        }
    }
}
