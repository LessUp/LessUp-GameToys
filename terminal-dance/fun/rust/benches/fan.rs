use criterion::{criterion_group, criterion_main, Criterion};

fn bench_frames(c: &mut Criterion) {
    let frames = fun::frame_set(false);
    c.bench_function("frame_advance", |b| {
        let mut idx = 0usize;
        b.iter(|| {
            idx = (idx + 1) % frames.len();
        });
    });
}

criterion_group!(benches, bench_frames);
criterion_main!(benches);
