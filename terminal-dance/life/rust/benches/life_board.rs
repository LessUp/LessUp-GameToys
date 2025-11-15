use criterion::{criterion_group, criterion_main, Criterion};

fn next_generation(c: &mut Criterion) {
    c.bench_function("board_next", |b| {
        let mut board = life::Board::empty(40, 20, false);
        b.iter(|| {
            board = board.next();
        });
    });
}

criterion_group!(benches, next_generation);
criterion_main!(benches);
