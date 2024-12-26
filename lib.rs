macro_rules! library {
    ($year:tt $path:tt $($day:tt),*) => {
        #[path = $path]
        pub mod $year {$(
            pub mod $day;
        )*}
    }
}

library!(year2015 "2015"
    day01
);
